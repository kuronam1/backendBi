package Store

import (
	"database/sql"
	"errors"
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
	var res []models.Group
	//GetLogger()

	stmt, err := g.store.DB.Prepare("SELECT * FROM groups")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	for rows.Next() {
		var group models.Group
		err = rows.Scan(&group.Id, &group.Name)
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

func (g *GroupRepository) GetGroupByName(name string) (*models.Group, error) {
	const op = "fc.groupRep.GetGroupByName"
	stmt, err := g.store.DB.Prepare("SELECT group_id, speciality, course FROM groups WHERE group_name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var gr = models.Group{
		Name: name,
	}
	err = stmt.QueryRow(name).Scan(&gr.Id, &gr.Speciality, &gr.Course)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, GroupNotRegistered
	case err != nil:
		return nil, err
	default:
		return &gr, nil
	}
}
