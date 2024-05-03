package UserRepository

import (
	"database/sql"
	"errors"
	"sbitnev_back/internal/database/models"
)

var (
	invalidUser = errors.New("user not registered")
)

type UserRepository struct {
	DB *sql.DB
}

func (u *UserRepository) GetUserByLogin(login string) (*models.User, error) {
	query := `SELECT * FROM users WHERE login = $1`

	DBrow := u.DB.QueryRow(query, login)
	if err := DBrow.Err(); err != nil {
		return nil, invalidUser
	}

	user := &models.User{}
	if err := DBrow.Scan(
		&user.UserId,
		&user.Login,
		&user.Password,
		&user.FullName,
		&user.Role,
		&user.GroupID,
	); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}

func (u *UserRepository) CreateUser(user *models.User) error {
	return nil
}

func (u *UserRepository) DeleteUser(user *models.User) error {
	return nil
}

func (u *UserRepository) UpdateUser(user *models.User) error {
	return nil
}
