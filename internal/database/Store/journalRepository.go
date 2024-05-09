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

func (j *JournalRepository) GetJournalByUserID(id int) (*models.Journal, error) {
	return nil, nil
}

func (j *JournalRepository) UpdateGrade(oldGrade, newGrade *models.Grade) error {
	const op = "fc.journalRep.UpdateGrade"
	stmt, err := j.store.DB.Prepare("UPDATE grades SET grade = $1, date = $2, comment = $3 WHERE grade_id = $4")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newGrade.Level, newGrade.Date, newGrade.Comment, oldGrade.GradeID)
	if err != nil {
		return err
	}

	return nil
}

func (j *JournalRepository) GetAdminJournal(groupName, disciplineName string) (*models.Journal, error) {
	const op = "fc.journalRep.GetAdminJournal"
	stmt, err := j.store.DB.Prepare(`
SELECT g.grade, g.date, g.comment, u.full_name
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
		Grades: make(map[string][]*models.Grade),
	}
	for rows.Next() {
		var (
			grade    *models.Grade
			fullName string
		)
		err = rows.Scan(
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
