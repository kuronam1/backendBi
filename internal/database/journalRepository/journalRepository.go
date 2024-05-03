package journalRepository

import (
	"database/sql"
	"sbitnev_back/internal/database/models"
)

type JournalRepository struct {
	DB *sql.DB
}

func GetJournalByUserID(id int) (*models.Journal, error) {
	return nil, nil
}
