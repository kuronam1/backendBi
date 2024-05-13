package Store

import (
	"database/sql"
	"errors"
	"fmt"
	"sbitnev_back/internal/database/models"
)

var (
	GroupNotRegistered = errors.New("group is not registered")
)

type GroupRepository struct {
	store *Storage
}

func (g *GroupRepository) GetAllGroups() ([]models.Group, error) {
	const op = "fc.Storage.GetAllGroups"

	stmt, err := g.store.DB.Prepare("SELECT group_id, group_name, number, speciality, course FROM groups")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Group
	for rows.Next() {
		group := models.Group{}
		err = rows.Scan(
			&group.Id,
			&group.Name,
			&group.Number,
			&group.Speciality,
			&group.Course)
		if err != nil {
			return nil, err
		}
		res = append(res, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GroupRepository) GetGroupByName(name interface{}) (*models.Group, error) {
	const op = "fc.groupRep.GetGroupByName"
	fmt.Println("i am here 2")
	stmt, err := g.store.DB.Prepare("SELECT group_id, group_name, number, speciality, course FROM groups WHERE group_name = $1")
	if err != nil {
		return nil, err
	}
	fmt.Println("i am here 1123123")
	defer stmt.Close()

	group := &models.Group{}
	err = stmt.QueryRow(name).Scan(&group.Id, &group.Name, &group.Number, &group.Speciality, &group.Course)
	fmt.Println("i am here 1123123")
	if err != nil {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, GroupNotRegistered
	} else if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *GroupRepository) GroupRegistration(group *models.Group) error {
	const op = "fc.groupRep.GroupRegistration"
	stmt, err := g.store.DB.Prepare("INSERT INTO groups (group_name, number, speciality, course) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING")
	if err != nil {
		return err
	}

	groupName := configurationGroupName(group.Speciality, group.Number, group.Course)

	_, err = stmt.Exec(groupName, group.Number, group.Speciality, group.Course)
	return err
}

func configurationGroupName(speciality string, number, course int) string {
	switch speciality {
	case "ЭВМ":
		return fmt.Sprintf("ЭВМ%d-%d", course, number)
	case "БИ":
		return fmt.Sprintf("БИ%d-%d", course, number)
	case "ПМ":
		return fmt.Sprintf("ПМ%d-%d", course, number)
	default:
		return fmt.Sprintf("БП%d-%d", course, number)
	}
}

func (g *GroupRepository) GetAllSpecialities() ([]string, error) {
	const op = "fc.groupRep.GetAllSpecialities"
	query := `SELECT speciality FROM disciplines`

	rows, err := g.store.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []string
	for rows.Next() {
		var speciality string
		err = rows.Scan(&speciality)
		if err != nil {
			return nil, err
		}
		res = append(res, speciality)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, err
}

func (g *GroupRepository) GetGroupByID(id int) (*models.Group, error) {
	const op = "fc.groupRep.GetGroupByID"
	stmt, err := g.store.DB.Prepare("SELECT group_name, speciality, number, course FROM groups WHERE group_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	group := &models.Group{}
	err = stmt.QueryRow(id).Scan(
		&group.Name,
		&group.Speciality,
		&group.Number,
		&group.Course)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (g *GroupRepository) GroupMembership(studentID int) (*models.Group, error) {
	stmt, err := g.store.DB.Prepare("SELECT g.group_id, group_name, number, speciality, course FROM groups g JOIN group_students gr ON g.group_id = gr.group_id WHERE student_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	group := &models.Group{}
	err = stmt.QueryRow(studentID).Scan(
		&group.Id,
		&group.Name,
		&group.Number,
		&group.Speciality,
		&group.Course)
	if err != nil {
		return nil, err
	}

	return group, nil
}
