package Store

import (
	"database/sql"
	"errors"
	"github.com/xuri/excelize/v2"
	"sbitnev_back/internal/database/models"
	"strconv"
	"time"
)

var (
	ErrNoLessons      = errors.New("no lessons")
	InternalServerErr = errors.New("internal server error")
)

type ScheduleRepository struct {
	store *Storage
}

func (s *ScheduleRepository) ScheduleRegister(filePath string) error {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rows, err := file.GetRows("Расписание")
	if err != nil {
		return err
	}

	var stmt *sql.Stmt
	defer stmt.Close()
	for _, row := range rows {
		var lesson models.Lesson
		lesson.GroupName = row[0]
		lesson.Time, err = time.Parse("1/2/06 15:04", row[1])
		if err != nil {
			return err
		}
		lesson.DisciplineName = row[2]
		lesson.Audience = row[3]
		lesson.Description = row[4]
		teacherName := row[5]
		lesson.LessonOrder, err = strconv.Atoi(row[6])
		if err != nil {
			return InternalServerErr
		}

		group, err := s.store.Groups().GetGroupByName(lesson.GroupName)
		if err != nil {
			return err
		}

		disID, err := s.store.Disciplines().RegisterDiscipline(teacherName,
			lesson.DisciplineName, group.Speciality, group.Course)
		if err != nil {
			return err
		}

		stmt, err = s.store.DB.Prepare("INSERT INTO lessons (group_id, time, discipline_id, audience, description, lesson_order) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(group.Id, lesson.Time, disID, lesson.Audience, lesson.Description)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ScheduleRepository) GetScheduleByGroupName(groupName string) (*models.Schedule, error) {
	const op = "fc.scheduleRep.GetScheduleByGroupName"
	schedule := make(map[time.Weekday][]models.Lesson)

	stmt, err := s.store.DB.Prepare(`
SELECT l.lesson_id,  g.group_name, l.time, d.discipline_name, l.audience, l.description, l.lesson_order, u.full_name
FROM lessons l
        JOIN disciplines d ON l.discipline_id = d.discipline_id
		JOIN users u ON u.user_id = d.teacher_id
		JOIN groups g ON l.group_id = g.group_id
WHERE g.group_name = $1
ORDER BY l.time`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lesson models.Lesson
		err = rows.Scan(
			&lesson.LessonId,
			&lesson.GroupName,
			&lesson.Time,
			&lesson.DisciplineName,
			&lesson.Description,
			&lesson.LessonOrder,
			&lesson.TeacherName)
		if err != nil {
			return nil, err
		}
		schedule[lesson.Time.Weekday()] = append(schedule[lesson.Time.Weekday()], lesson) // Sunday = 0, ... !
	}

	return &models.Schedule{Lessons: schedule}, err
}

func (s *ScheduleRepository) GetScheduleByTeacherName(teacherName string) (*models.Schedule, error) {
	const op = "fc.scheduleRep.GetScheduleByTeacherName"
	schedule := make(map[time.Weekday][]models.Lesson)

	stmt, err := s.store.DB.Prepare(`
SELECT g.group_name, l.time, d.discipline_name, l.audience, l.description, l.lesson_order
FROM lessons l
		JOIN disciplines d ON l.description = d.discipline_id
    	JOIN groups g ON l.group_id = g.group_id
		JOIN users u ON u.user_id = d.teacher_id
WHERE u.full_name = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(teacherName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lesson models.Lesson
		err = rows.Scan(
			&lesson.GroupName,
			&lesson.Time,
			&lesson.DisciplineName,
			&lesson.Audience,
			&lesson.Description,
			&lesson.LessonOrder)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoLessons
		case err != nil:
			return nil, err
		}
		schedule[lesson.Time.Weekday()] = append(schedule[lesson.Time.Weekday()], lesson) // Sunday = 0, ... !
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &models.Schedule{
		Lessons: schedule,
	}, err
}

func (s *ScheduleRepository) GetScheduleByID(id int) (*models.Schedule, error) {
	const op = "fc.scheduleRep.GetScheduleByID"
	stmt, err := s.store.DB.Prepare(`
SELECT g.group_name, l.time, d.discipline_name, l.audience, l.description, u.full_name AS teacher_name, l.lesson_order
FROM users u
         JOIN disciplines d ON d.teacher_id = u.user_id
         JOIN lessons l ON l.discipline_id = d.discipline_id
         JOIN groups g ON l.group_id = g.group_id
         JOIN group_students gs ON g.group_id = gs.group_id
WHERE gs.student_id = $1
ORDER BY l.time`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schedule := &models.Schedule{
		Lessons: make(map[time.Weekday][]models.Lesson),
	}
	for rows.Next() {
		var lesson models.Lesson
		err = rows.Scan(
			&lesson.GroupName,
			&lesson.Time,
			&lesson.DisciplineName,
			&lesson.Audience,
			&lesson.Description,
			&lesson.TeacherName,
			&lesson.LessonOrder)
		if err != nil {
			return nil, err
		}
		schedule.Lessons[lesson.Time.Weekday()] = append(schedule.Lessons[lesson.Time.Weekday()], lesson)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedule, nil
}
