package Store

import (
	"database/sql"
	"errors"
	"fmt"
	"sbitnev_back/internal/database/models"
	"slices"
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
	stmt, err := g.store.DB.Prepare("SELECT group_id, group_name, number, speciality, course FROM groups WHERE group_name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	group := &models.Group{}
	err = stmt.QueryRow(name).Scan(&group.Id, &group.Name, &group.Number, &group.Speciality, &group.Course)
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
	case "Специальные машины и устройства":
		return fmt.Sprintf("СЦУ%d%d", course, number)
	case "Автоматические системы управления":
		return fmt.Sprintf("АСУ%d%d", course, number)
	case "Компьютерные системы и комплексы":
		return fmt.Sprintf("КСК%d%d", course, number)
	case "Контроль работы измерительных приборов":
		return fmt.Sprintf("КРИП%d%d", course, number)
	case "Экономика и бухгалтерский учет":
		return fmt.Sprintf("ЭБУ%d%d", course, number)
	case "Технология металлообрабатывающего производства":
		return fmt.Sprintf("ТМП%d%d", course, number)
	case "Технология машиностроения":
		return fmt.Sprintf("ТМ%d%d", course, number)
	case "Информационные системы и программирование":
		return fmt.Sprintf("ИСП%d%d", course, number)
	case "Техническая эксплуатация летательных аппаратов и двигателей":
		return fmt.Sprintf("ТЭЛАД%d%d", course, number)
	default:
		return fmt.Sprintf("ТЭЭПНК%d%d", course, number)
	}
}

func (g *GroupRepository) GetAllSpecialities() ([]string, error) {
	const op = "fc.groupRep.GetAllSpecialities"
	query := `SELECT speciality_name FROM specialities`

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

func (g *GroupRepository) GetAllTeachersGroups(teacherID int) ([]string, error) {
	stmt, err := g.store.DB.Prepare(`SELECT group_name FROM groups g
    JOIN disciplines d ON d.course = g.course AND d.speciality = g.speciality
WHERE teacher_id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var groupName string
		err := rows.Scan(&groupName)
		if err != nil {
			return nil, err
		}
		if !slices.Contains(result, groupName) {
			result = append(result, groupName)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (g *GroupRepository) UpdateGroupsCourse() error {
	stmt, err := g.store.DB.Prepare("UPDATE groups SET course = course + 1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	query := `UPDATE groups SET group_name = $1 WHERE group_id = $2`

	groups, err := g.store.Groups().GetAllGroups()
	if err != nil {
		return err
	}

	for _, group := range groups {
		groupName := configurationGroupName(group.Speciality, group.Number, group.Course)
		_, err = g.store.DB.Exec(query, groupName, group.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *GroupRepository) DeleteGroup(group models.Group) error {
	stmt, err := g.store.DB.Prepare(`
	DELETE FROM groups 
	WHERE group_name = $1`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(group.Name)
	return err
}

func (g *GroupRepository) DeleteGroupLink(groupID int) error {
	stmt, err := g.store.DB.Prepare("DELETE FROM group_students WHERE group_id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&groupID)
	return err
}
