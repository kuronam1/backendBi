package Store

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"sbitnev_back/internal/database/models"
	"strconv"
	"time"
)

var (
	ErrNoLessons      = errors.New("no lessons")
	InternalServerErr = errors.New("internal server error")
)

const timePayload = "02-01-2006"

type ScheduleRepository struct {
	store *Storage
}

func (s *ScheduleRepository) ScheduleRegister(filePath, groupName string) error {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rows, err := file.GetRows("Расписание")
	if err != nil {
		return err
	}

	stmt, err := s.store.DB.Prepare(`
				INSERT INTO lessons (group_id, time, discipline_id, teacher_id, audience, description, lesson_order)
				VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, row := range rows {
		lesson := models.Lesson{}
		lesson.GroupName = groupName
		lesson.Time, err = time.Parse(time.DateOnly, row[0])
		if err != nil {
			return err
		}
		lesson.DisciplineName = row[1]
		lesson.Audience = row[2]
		lesson.Description = row[3]
		teacherName := row[4]
		lesson.LessonOrder, err = strconv.Atoi(row[5])
		if err != nil {
			return InternalServerErr
		}

		group, err := s.store.Groups().GetGroupByName(lesson.GroupName)
		if err != nil {
			return fmt.Errorf("group not found")
		}

		discipline, err := s.store.Disciplines().GetDisciplineByName(lesson.DisciplineName)
		if err != nil {
			return err
		}

		teacher, err := s.store.User().GetUserByName(teacherName)
		if err != nil {
			return fmt.Errorf("teacher not found")
		}

		_, err = stmt.Exec(group.Id, lesson.Time,
			discipline.DisciplineID, teacher.UserID, lesson.Audience,
			lesson.Description, lesson.LessonOrder)
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

	stmt, err := s.store.DB.Prepare("SELECT lesson_id, time, discipline_id, teacher_id, audience, description, subject, homework, lesson_order FROM lessons WHERE group_id = $1")
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
			&parse.Subject,
			&parse.HomeWork,
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
			Subject:        parse.Subject.String,
			HomeWork:       parse.HomeWork.String,
			LessonOrder:    parse.LessonOrder,
		}
		schedule.Lessons[lesson.Time.Weekday()] = append(schedule.Lessons[lesson.Time.Weekday()], lesson) // Sunday = 0, ... !
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

	return result, err
}

func (s *ScheduleRepository) GetScheduleByTeacherName(teacherName string) ([]map[string][]models.Lesson, error) {
	const op = "fc.scheduleRep.GetScheduleByTeacherName"

	teacher, err := s.store.User().GetUserByName(teacherName)
	if err != nil {
		return nil, err
	}

	stmt, err := s.store.DB.Prepare("SELECT lesson_id, group_id, time, discipline_id, audience, description, subject, homework, lesson_order FROM lessons WHERE teacher_id = $1")
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
			&parse.Subject,
			&parse.HomeWork,
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
			Subject:        parse.Subject.String,
			HomeWork:       parse.HomeWork.String,
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

	stmt, err := s.store.DB.Prepare("SELECT lesson_id, group_id, time, discipline_id, audience, description, subject, homework, lesson_order FROM lessons WHERE teacher_id = $1")
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
			&parse.Subject,
			&parse.HomeWork,
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
			Subject:        parse.Subject.String,
			HomeWork:       parse.HomeWork.String,
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
		tm := t.Format(timePayload)
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

func (s *ScheduleRepository) UpdateHomeWork(homeWork string, id int) error {
	stmt, err := s.store.DB.Prepare("UPDATE lessons SET homework = $1 WHERE lesson_id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(homeWork, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRepository) UpdateSubject(subject string, id int) error {
	stmt, err := s.store.DB.Prepare("UPDATE lessons SET subject = $1 WHERE lesson_id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(subject, id)
	return err
}

func (s *ScheduleRepository) UpdateSubjectAndHomeWork(subject, homeWork string, id int) error {
	stmt, err := s.store.DB.Prepare("UPDATE lessons SET subject = $1, homework = $2 WHERE lesson_id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(subject, homeWork, id)
	return err
}

func (s *ScheduleRepository) ClearSchedule() error {
	query := `
		DROP TABLE lessons;
		CREATE TABLE IF NOT EXISTS lessons (
		   					lesson_id SERIAL PRIMARY KEY,
            				group_id INTEGER NOT NULL REFERENCES groups(group_id),
                            time DATE NOT NULL,
                            discipline_id INTEGER NOT NULL REFERENCES disciplines(discipline_id),
                            teacher_id INTEGER NOT NULL REFERENCES users(user_id),
                            audience VARCHAR(10) NOT NULL,
                            description VARCHAR NOT NULL,
                            subject VARCHAR DEFAULT NULL,
                            homework VARCHAR DEFAULT NULL,
                            lesson_order INTEGER NOT NULL
		)`

	_, err := s.store.DB.Exec(query)
	return err
}
