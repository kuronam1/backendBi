package Store

import (
	"errors"
	"fmt"
	"sbitnev_back/internal/database/models"
	"time"
)

var (
	NotRegistered = errors.New("discipline is not registered")
	NoGrades      = errors.New("no grades")
)

type JournalRepository struct {
	store *Storage
}

func (j *JournalRepository) GetJournalByStudentID(id int) ([]map[string][]models.Grade, error) {
	const op = "fc.journalRep.UpdateGrade"
	stmt, err := j.store.DB.Prepare(`
SELECT d.discipline_name, g.grade, g.date, g.comment FROM grades g
        JOIN disciplines d ON g.discipline_id = d.discipline_id
        JOIN users u ON g.student_id = u.user_id
WHERE u.user_id = $1
ORDER BY g.date`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	journal := &models.Journal{
		Grades: make(map[string][]models.Grade),
	}
	for rows.Next() {
		var (
			grade   models.Grade
			disName string
		)
		err := rows.Scan(
			&disName,
			&grade.Level,
			&grade.Date,
			&grade.Comment)
		if err != nil {
			return nil, err
		}
		journal.Grades[disName] = append(journal.Grades[disName], grade)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	group, err := j.store.Groups().GroupMembership(id)
	if err != nil {
		return nil, err
	}

	GroupDis, err := j.store.Disciplines().GetGroupDisciplines(group.Name)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(GroupDis); i++ {
		journal.Headers = append(journal.Headers, GroupDis[i].DisciplineName)
	}

	result := make([]map[string][]models.Grade, len(journal.Headers))
	for i := 0; i < len(journal.Headers); i++ {
		mp := make(map[string][]models.Grade)
		var disciplineGrades []models.Grade
		mp[journal.Headers[i]] = disciplineGrades
		_, inMap := journal.Grades[journal.Headers[i]]
		if inMap {
			mp[journal.Headers[i]] = journal.Grades[journal.Headers[i]]
		}
		result[i] = mp
	}

	return result, nil
}

func (j *JournalRepository) UpdateGrade(oldGrade, newGrade *models.Grade) error {
	const op = "fc.journalRep.UpdateGrade"
	stmt, err := j.store.DB.Prepare("UPDATE grades SET grade = $1, date = $2, comment = $3 WHERE grade_id = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newGrade.Level, newGrade.Date, newGrade.Comment, oldGrade.GradeID)
	if err != nil {
		return err
	}

	return nil
}

func (j *JournalRepository) GetGroupJournalByDiscipline(groupName, disciplineName string) ([]map[string][]models.Grade, error) {
	const op = "fc.journalRep.GetAdminJournal"
	stmt, err := j.store.DB.Prepare(`
SELECT g.grade_id, g.discipline_id, g.grade, g.date, g.comment, u.full_name
FROM grades g
			JOIN users u ON g.student_id = u.user_id
			JOIN disciplines d ON d.discipline_id = g.discipline_id
			JOIN groups gr ON d.speciality = gr.speciality AND d.course = gr.course
WHERE gr.group_name = $1 AND d.discipline_name = $2
ORDER BY u.full_name`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupName, disciplineName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	journal := &models.Journal{
		Grades: make(map[string][]models.Grade),
	}

	for rows.Next() {
		var (
			grade    models.Grade
			fullName string
		)
		err = rows.Scan(
			&grade.GradeID,
			&grade.DisciplineID,
			&grade.Level,
			&grade.Date,
			&grade.Comment,
			&fullName)
		if err != nil {
			return nil, err
		}
		fmt.Println(grade)
		journal.Grades[fullName] = append(journal.Grades[fullName], grade)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	journal.Headers, err = j.store.User().GetUserNamesByGroup(groupName)
	if err != nil {
		return nil, err
	}

	lessons, err := j.store.Schedule().GetAllGroupsLessonsOneDis(groupName, disciplineName)
	if err != nil {
		return nil, err
	}

	result := make([]map[string][]models.Grade, len(journal.Headers))
	for i := 0; i < len(journal.Headers); i++ {
		mp := make(map[string][]models.Grade)
		studentGrades := make([]models.Grade, len(lessons))
		for j := 0; j < len(lessons); j++ {
			grades, inMap := journal.Grades[journal.Headers[i]]
			if inMap {
				for _, value := range grades {
					if value.Date.Format(time.DateOnly) == lessons[j] {
						studentGrades[j] = value
					}
				}
			}
		}
		mp[journal.Headers[i]] = studentGrades
		result[i] = mp
	}

	return result, nil
}

func (j *JournalRepository) CreateGrade(grade *models.Grade) error {
	const op = "fc.journalRep.GetAdminJournal"
	stmt, err := j.store.DB.Prepare("INSERT INTO grades (student_id, discipline_id, grade, date, comment) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(grade.StudentID, grade.DisciplineID, grade.Level, grade.Date, grade.Comment)

	return err
}
