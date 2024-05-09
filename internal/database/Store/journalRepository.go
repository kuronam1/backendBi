package Store

import (
	"errors"
	"sbitnev_back/internal/database/models"
)

var (
	NotRegistered = errors.New("discipline is not registered")
)

type JournalRepository struct {
	store *Storage
}

func (j *JournalRepository) GetJournalByUserID(id int) (*models.Journal, error) {
	return nil, nil
}

func (j *JournalRepository) UpdateGrade(oldGrade, newGrade *models.Grade) error {
	const op = "fc.journalRep.UpdateGrade"
	stmt, err := j.store.DB.Prepare(`UPDATE grades SET grade = $1, time = $2, comment = $3 WHERE grade_id = $4`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newGrade.Level, newGrade.Date, newGrade.Comment, oldGrade.GradeID)
	if err != nil {
		return err
	}

	return nil
}
