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

		stmt, err = s.store.DB.Prepare("INSERT INTO lessons (group_id, time, discipline_id, teacher_id, audience, description, lesson_order) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(group.Id, lesson.Time, disID, lesson.Audience, lesson.Description) // добавить teacherID
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ScheduleRepository) GetScheduleByGroupName(groupName interface{}) ([]map[string][]models.Lesson, error) {
	const op = "fc.scheduleRep.GetScheduleByGroupName"
	schedule := models.Schedule{
		Lessons: make(map[time.Weekday][]models.Lesson),
	}

	group, err := s.store.Groups().GetGroupByName(groupName)
	if err != nil {
		return nil, err
	}

	stmt, err := s.store.DB.Prepare("SELECT lesson_id, time, discipline_id, teacher_id, audience, description, lesson_order FROM lessons WHERE group_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(group.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		parse := &models.ParseLessons{
			GroupID: group.Id,
		}
		err = rows.Scan(
			&parse.LessonID,
			&parse.Time,
			&parse.DisciplineID,
			&parse.TeacherID,
			&parse.Audience,
			&parse.Description,
			&parse.LessonOrder)
		if err != nil {
			return nil, err
		}

		discipline, err := s.store.Disciplines().GetDisciplineByID(parse.DisciplineID)
		if err != nil {
			return nil, err
		}

		teacher, err := s.store.User().GetUserByID(parse.TeacherID)
		if err != nil {
			return nil, err
		}

		lesson := models.Lesson{
			LessonId:       parse.LessonID,
			GroupName:      groupName.(string),
			Time:           parse.Time,
			DisciplineName: discipline.DisciplineName,
			Audience:       parse.Audience,
			Description:    parse.Description,
			TeacherName:    teacher.FullName,
			LessonOrder:    parse.LessonOrder,
		}
		schedule.Lessons[lesson.Time.Weekday()] = append(schedule.Lessons[lesson.Time.Weekday()], lesson) // Sunday = 0, ... !
	}

	schedule.Headers = []string{"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"}

	result := make([]map[string][]models.Lesson, len(schedule.Headers))
	for i := time.Weekday(0); i <= time.Weekday(6); i++ {
		mp := make(map[string][]models.Lesson)
		lessons := make([]models.Lesson, 6)
		for j := 1; j <= 6; j++ {
			lesson, inMap := schedule.Lessons[i]
			if inMap {
				for _, l := range lesson {
					if l.LessonOrder == j {
						lessons[j] = l
					}
				}
			}
		}
		mp[schedule.Headers[i]] = lessons
		result[i] = mp
	}

	return result, err
}

func (s *ScheduleRepository) GetScheduleByTeacherName(teacherName string) ([]map[string][]models.Lesson, error) {
	const op = "fc.scheduleRep.GetScheduleByTeacherName"

	teacher, err := s.store.User().GetUserByName(teacherName)
	if err != nil {
		return nil, err
	}

	stmt, err := s.store.DB.Prepare("SELECT lesson_id, group_id, time, discipline_id, audience, description, lesson_order FROM lessons WHERE teacher_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(teacher.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schedule := &models.Schedule{
		Lessons: make(map[time.Weekday][]models.Lesson),
	}
	for rows.Next() {
		parse := models.ParseLessons{}
		err = rows.Scan(
			&parse.LessonID,
			&parse.GroupID,
			&parse.Time,
			&parse.DisciplineID,
			&parse.Audience,
			&parse.Description,
			&parse.LessonOrder)
		if err != nil {
			return nil, err
		}

		discipline, err := s.store.Disciplines().GetDisciplineByID(parse.DisciplineID)
		if err != nil {
			return nil, err
		}

		group, err := s.store.Groups().GetGroupByID(parse.GroupID)
		if err != nil {
			return nil, err
		}

		lesson := models.Lesson{
			LessonId:       parse.LessonID,
			GroupName:      group.Name,
			Time:           parse.Time,
			DisciplineName: discipline.DisciplineName,
			Audience:       parse.Audience,
			Description:    parse.Description,
			TeacherName:    teacher.FullName,
			LessonOrder:    parse.LessonOrder,
		}
		schedule.Lessons[lesson.Time.Weekday()] = append(schedule.Lessons[lesson.Time.Weekday()], lesson)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	schedule.Headers = []string{"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"}

	result := make([]map[string][]models.Lesson, len(schedule.Headers))
	for i := time.Weekday(0); i <= time.Weekday(6); i++ {
		mp := make(map[string][]models.Lesson)
		lessons := make([]models.Lesson, 5)
		for j := 0; j < 5; j++ {
			lesson, inMap := schedule.Lessons[i]
			if inMap {
				for _, l := range lesson {
					if l.LessonOrder == j {
						lessons[j] = l
					}
				}
			}
		}
		mp[schedule.Headers[i]] = lessons
		result[i] = mp
	}

	return result, nil
}

func (s *ScheduleRepository) GetScheduleByTeacherID(id int) ([]map[string][]models.Lesson, error) {
	const op = "fc.scheduleRep.GetScheduleByTeacherID"
	teacher, err := s.store.User().GetUserByID(id)
	if err != nil {
		return nil, err
	}

	stmt, err := s.store.DB.Prepare("SELECT lesson_id, group_id, time, discipline_id, audience, description, lesson_order FROM lessons WHERE teacher_id = $1")
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
		parse := models.ParseLessons{}
		err = rows.Scan(
			&parse.LessonID,
			&parse.GroupID,
			&parse.Time,
			&parse.DisciplineID,
			&parse.Audience,
			&parse.Description,
			&parse.LessonOrder)
		if err != nil {
			return nil, err
		}

		discipline, err := s.store.Disciplines().GetDisciplineByID(parse.DisciplineID)
		if err != nil {
			return nil, err
		}

		group, err := s.store.Groups().GetGroupByID(parse.GroupID)
		if err != nil {
			return nil, err
		}

		lesson := models.Lesson{
			LessonId:       parse.LessonID,
			GroupName:      group.Name,
			Time:           parse.Time,
			DisciplineName: discipline.DisciplineName,
			Audience:       parse.Audience,
			Description:    parse.Description,
			TeacherName:    teacher.FullName,
			LessonOrder:    parse.LessonOrder,
		}
		schedule.Lessons[lesson.Time.Weekday()] = append(schedule.Lessons[lesson.Time.Weekday()], lesson)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	schedule.Headers = []string{"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"}

	result := make([]map[string][]models.Lesson, len(schedule.Headers))
	for i := time.Weekday(0); i <= time.Weekday(6); i++ {
		mp := make(map[string][]models.Lesson)
		lessons := make([]models.Lesson, 5)
		for j := 0; j < 5; j++ {
			lesson, inMap := schedule.Lessons[i]
			if inMap {
				for _, l := range lesson {
					if l.LessonOrder == j {
						lessons[j] = l
					}
				}
			}
		}
		mp[schedule.Headers[i]] = lessons
		result[i] = mp
	}

	return result, nil
}

func (s *ScheduleRepository) GetAllGroupsLessonsOneDis(groupName, disciplineName interface{}) ([]string, error) {
	group, err := s.store.Groups().GetGroupByName(groupName)
	if err != nil {
		return nil, err
	}

	discipline, err := s.store.Disciplines().GetDisciplineByName(disciplineName)

	stmt, err := s.store.DB.Prepare("SELECT time FROM lessons WHERE group_id = $1 AND discipline_id = $2 ORDER BY time")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(group.Id, discipline.DisciplineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []string
	for rows.Next() {
		var t time.Time
		err = rows.Scan(&t)
		tm := t.Format(time.DateOnly)
		if err != nil {
			return nil, err
		}
		res = append(res, tm)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
