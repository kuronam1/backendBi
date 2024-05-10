package Store

import (
	"errors"
	"sbitnev_back/internal/database/models"
)

var (
	NotRegistered = errors.New("discipline is not registered")
	NoGrades      = errors.New("no grades")
)

type JournalRepository struct {
	store *Storage
}

func (j *JournalRepository) GetJournalByStudentID(id int) (*models.Journal, error) {
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

	return journal, nil

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

func (j *JournalRepository) GetGroupJournalByDiscipline(groupName, disciplineName string) (*models.Journal, error) {
	const op = "fc.journalRep.GetAdminJournal"
	stmt, err := j.store.DB.Prepare(`
SELECT g.grade_id, g.grade, g.date, g.comment, u.full_name
FROM grades g
			JOIN users u ON g.student_id = u.user_id
			JOIN disciplines d ON d.discipline_id = g.discipline_id
			JOIN groups gr ON d.speciality = gr.speciality AND d.course = gr.course
WHERE gr.group_name = $1 AND d.discipline_name = $2`)
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
			&grade.Level,
			&grade.Date,
			&grade.Comment,
			&fullName)
		if err != nil {
			return nil, err
		}
		journal.Grades[fullName] = append(journal.Grades[fullName], grade)
	}

	return journal, nil
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
