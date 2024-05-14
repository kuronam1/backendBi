package Store

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
	store *Storage
}

func (u *UserRepository) GetUserByLogin(login string) (*models.User, error) {
	const op = "fc.userRep.GetUserByLogin"

	stmt, err := u.store.DB.Prepare("SELECT * FROM users WHERE login = $1")
	if err != nil {
		return nil, internalServerErr
	}
	defer stmt.Close()

	user := &models.User{}
	err = stmt.QueryRow(login).Scan(
		&user.UserID,
		&user.Login,
		&user.Password,
		&user.FullName,
		&user.Role,
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
	const op = "fc.userRep.GetUserByID"
	stmt, err := u.store.DB.Prepare("SELECT login, password, full_name, role FROM users WHERE user_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &models.User{}
	err = stmt.QueryRow(id).Scan(
		&user.Login,
		&user.Password,
		&user.FullName,
		&user.Role)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, invalidUser
	case err != nil:
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUserByName(name string) (*models.User, error) {
	const op = "fc.userRep.GetUserByName"
	stmt, err := u.store.DB.Prepare("SELECT user_id, login, password, role FROM users WHERE full_name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user = models.User{
		FullName: name,
	}
	err = stmt.QueryRow(name).Scan(
		&user.UserID,
		&user.Login,
		&user.Password,
		&user.Role,
	)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, invalidUser
	case err != nil:
		return nil, internalServerErr
	default:
		return &user, nil
	}
}

func (u *UserRepository) CreateUser(user *models.User) (int, error) {
	const op = "fc.userRep.CreateUser"
	stmt, err := u.store.DB.Prepare("INSERT INTO users (login, password, full_name, role) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Login, user.Password, user.FullName, user.Role)
	if err != nil {
		return 0, err
	}

	userData, err := u.GetUserByLogin(user.Login)
	if err != nil {
		return 0, err
	}

	return userData.UserID, nil
}

/*func (u *UserRepository) DeleteUser(user *models.User) error {
	return nil
}

func (u *UserRepository) UpdateUser(user *models.User) error {
	return nil
}*/

func (u *UserRepository) CreateGroupUserLink(userID int, groupName string) error {
	const op = "fc.userRep.CreateUserLink"

	selectStmt, err := u.store.DB.Prepare("SELECT group_id FROM groups WHERE group_name = $1")
	if err != nil {
		return err
	}
	defer selectStmt.Close()

	insertStmt, err := u.store.DB.Prepare("INSERT INTO group_students VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer insertStmt.Close()

	var groupID int
	if err := selectStmt.QueryRow(groupName).Scan(&groupID); err != nil {
		return err
	}

	_, err = insertStmt.Exec(groupID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) CreateTeacherDisciplineLink(teacherID int, disciplineName string) error {
	stmt, err := u.store.DB.Prepare("UPDATE disciplines SET teacher_id = $1 WHERE discipline_name = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(teacherID, disciplineName)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) CreateParentStudentLink(parentID int, studentID int) error {
	stmt, err := u.store.DB.Prepare("INSERT INTO parent_students VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(parentID, studentID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetAllTeachers() ([]models.User, error) {
	const op = "fc.userRep.GetAllTeachers"

	stmt, err := u.store.DB.Prepare("SELECT user_id, login, password, full_name, role FROM users WHERE role = 'teacher'")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.User
	for rows.Next() {
		teacher := models.User{}
		err := rows.Scan(
			&teacher.UserID,
			&teacher.Login,
			&teacher.Password,
			&teacher.FullName,
			&teacher.Role,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, teacher)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserRepository) GetUserNamesByGroup(groupName string) ([]string, error) {
	group, err := u.store.Groups().GetGroupByName(groupName)
	if err != nil {
		return nil, err
	}

	stmt, err := u.store.DB.Prepare("SELECT u.full_name FROM users u JOIN group_students gs ON u.user_id = gs.student_id WHERE group_id = $1 ORDER BY u.full_name")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(group.Id)
	if err != nil {
		return nil, err
	}

	var studentNames []string
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		studentNames = append(studentNames, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return studentNames, nil
}

func (u *UserRepository) GetStudentIDByParentID(parentID int) (int, error) {
	stmt, err := u.store.DB.Prepare("SELECT student_id FROM parent_students WHERE parent_id = $1")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var studentID int
	if err = stmt.QueryRow(parentID).Scan(studentID); err != nil {
		return 0, err
	}

	return studentID, nil

}
