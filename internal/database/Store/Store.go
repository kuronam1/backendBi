package Store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"os"
	"sbitnev_back/internal/config"
	"sbitnev_back/internal/database/UserRepository"
	"sbitnev_back/internal/database/journalRepository"
	"sbitnev_back/internal/database/scheduleRepository"
)

const (
	initFileName = "init.sql"
)

//const  = "postgres://user:Rtyuehe1223@localhost:5432/MyDB"

type Storage struct {
	DB              *sql.DB
	UserMethods     *UserRepository.UserRepository
	JournalMethods  *journalRepository.JournalRepository
	ScheduleMethods *scheduleRepository.ScheduleRepository
}

func InitStorage(c *config.Config) (*Storage, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.PG.Username, c.PG.Password, c.PG.Host, c.PG.Port, c.PG.DataBasename)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{
		DB: db,
	}, nil
}

func (s *Storage) CloseStore() {
	_ = s.DB.Close()
}

func (s *Storage) PrepareTables() error {
	file, err := os.Open(initFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	query, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	_, err = s.DB.Exec(string(query))
	return err
}

func (s *Storage) User() *UserRepository.UserRepository {
	if s.UserMethods != nil {
		return s.UserMethods
	}

	s.UserMethods = &UserRepository.UserRepository{
		DB: s.DB,
	}
	return s.UserMethods
}

func (s *Storage) Journal() *journalRepository.JournalRepository {
	if s.JournalMethods != nil {
		return s.JournalMethods
	}

	s.JournalMethods = &journalRepository.JournalRepository{
		DB: s.DB,
	}
	return s.JournalMethods
}

func (s *Storage) Schedule() *scheduleRepository.ScheduleRepository {
	if s.ScheduleMethods != nil {
		return s.ScheduleMethods
	}

	s.ScheduleMethods = &scheduleRepository.ScheduleRepository{
		DB: s.DB,
	}
	return s.ScheduleMethods
}

//Store.Repository(user).GetUserById(id)
