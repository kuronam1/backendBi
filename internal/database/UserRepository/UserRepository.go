package UserRepository

import (
	"database/sql"
	"errors"
	"fmt"
	"sbitnev_back/internal/database/models"
)

var (
	internalServerErr = errors.New("internal server error")
	invalidUser       = errors.New("user not registered")
)

type UserRepository struct {
	DB *sql.DB
}

func (u *UserRepository) GetUserByLogin(login string) (*models.User, error) {
	const op = "userRep.GetUserByLogin"

	stmt, err := u.DB.Prepare("SELECT * FROM users WHERE login = $1")
	if err != nil {
		return nil, internalServerErr
	}
	defer stmt.Close()

	user := &models.User{}
	err = stmt.QueryRow(login).Scan(
		&user.UserId,
		&user.Login,
		&user.Password,
		&user.FullName,
		&user.Role,
		&user.GroupID,
	)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, invalidUser
	case err != nil:
		return nil, fmt.Errorf("[%s]: %w", op, err)
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
