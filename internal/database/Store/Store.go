package Store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"io"
	"os"
	"sbitnev_back/internal/config"
)

const (
	initFileName = "./internal/database/Store/init.sql"
)

type Storage struct {
	DB                *sql.DB
	UserMethods       *UserRepository
	JournalMethods    *JournalRepository
	ScheduleMethods   *ScheduleRepository
	GroupMethods      *GroupRepository
	DisciplineMethods *DisciplineRepository
}

func InitStorage(c *config.Config) (*Storage, error) {
	db, err := sql.Open("postgres", c.PgUrl)
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

func (s *Storage) User() *UserRepository {
	if s.UserMethods != nil {
		return s.UserMethods
	}

	s.UserMethods = &UserRepository{
		store: s,
	}
	return s.UserMethods
}

func (s *Storage) Journal() *JournalRepository {
	if s.JournalMethods != nil {
		return s.JournalMethods
	}

	s.JournalMethods = &JournalRepository{
		store: s,
	}
	return s.JournalMethods
}

func (s *Storage) Schedule() *ScheduleRepository {
	if s.ScheduleMethods != nil {
		return s.ScheduleMethods
	}

	s.ScheduleMethods = &ScheduleRepository{
		store: s,
	}
	return s.ScheduleMethods
}

func (s *Storage) Groups() *GroupRepository {
	if s.GroupMethods != nil {
		return s.GroupMethods
	}

	s.GroupMethods = &GroupRepository{
		store: s,
	}
	return s.GroupMethods
}

func (s *Storage) Disciplines() *DisciplineRepository {
	if s.DisciplineMethods != nil {
		return s.DisciplineMethods
	}

	s.DisciplineMethods = &DisciplineRepository{
		store: s,
	}
	return s.DisciplineMethods
}
