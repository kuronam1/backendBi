package Store

import (
	"database/sql"
	"errors"
	"sbitnev_back/internal/database/models"
)

type DisciplineRepository struct {
	store *Storage
}

func (d *DisciplineRepository) GetDisciplineByName(name interface{}) (*models.Discipline, error) {
	const op = "fc.journalRep.GetDisciplineByName"
	stmt, err := d.store.DB.Prepare("SELECT discipline_id, teacher_id, speciality, course FROM disciplines WHERE discipline_name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var dis = models.Discipline{
		DisciplineName: name.(string),
	}
	err = stmt.QueryRow(name).Scan(
		&dis.DisciplineID,
		&dis.TeacherID,
		&dis.Speciality,
		&dis.Course,
	)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, NotRegistered
	case err != nil:
		return nil, internalServerErr
	default:
		return &dis, nil
	}
}

func (d *DisciplineRepository) RegisterDiscipline(teacherName, disciplineName, speciality string, course int) error {
	const op = "fc.discRep.RegisterDiscipline"
	stmt, err := d.store.DB.Prepare(`
		INSERT INTO disciplines (teacher_id, discipline_name, speciality, course)
		VALUES ($1, $2, $3, $4) RETURNING discipline_id`)
	if err != nil {
		return err
	}

	teacher, err := d.store.User().GetUserByName(teacherName)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(teacher.UserID, disciplineName, speciality, course)
	if err != nil {
		return err
	}

	/*id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}*/

	return nil
}

func (d *DisciplineRepository) GetAllDisciplines() ([]*models.Discipline, error) {
	const op = "fc.discRep.GetAllDisciplines"
	stmt, err := d.store.DB.Prepare("SELECT teacher_id, discipline_name, speciality, course FROM disciplines")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*models.Discipline
	for rows.Next() {
		discipline := &models.Discipline{}
		err = rows.Scan(
			&discipline.TeacherID,
			&discipline.DisciplineName,
			&discipline.Speciality,
			&discipline.Course)
		if err != nil {
			return nil, err
		}
		res = append(res, discipline)
	}

	return res, nil
}

func (d *DisciplineRepository) GetGroupDisciplines(groupName string) ([]models.Discipline, error) {
	group, err := d.store.Groups().GetGroupByName(groupName)
	if err != nil {
		return nil, err
	}

	stmt, err := d.store.DB.Prepare("SELECT discipline_id, teacher_id, discipline_name, speciality, course FROM disciplines WHERE speciality = $1 AND course = $2 ORDER BY discipline_name")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(group.Speciality, group.Course)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var disciplines []models.Discipline
	for rows.Next() {
		var discipline models.Discipline
		err = rows.Scan(
			&discipline.DisciplineID,
			&discipline.TeacherID,
			&discipline.DisciplineName,
			&discipline.Speciality,
			&discipline.Course)
		if err != nil {
			return nil, err
		}
		disciplines = append(disciplines, discipline)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return disciplines, nil
}

func (d *DisciplineRepository) GetDisciplinesByTeacherId(id int) ([]*models.Discipline, error) {
	stmt, err := d.store.DB.Prepare("SELECT discipline_id, discipline_name, speciality, course FROM disciplines WHERE teacher_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*models.Discipline
	for rows.Next() {
		discipline := &models.Discipline{}
		err := rows.Scan(
			&discipline.DisciplineID,
			&discipline.DisciplineName,
			&discipline.Speciality,
			&discipline.Course,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, discipline)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (d *DisciplineRepository) GetDisciplineByID(id int) (*models.Discipline, error) {
	stmt, err := d.store.DB.Prepare("SELECT teacher_id, discipline_name, speciality, course FROM disciplines WHERE disciplines.discipline_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	discipline := &models.Discipline{}
	err = stmt.QueryRow(id).Scan(
		&discipline.TeacherID,
		&discipline.DisciplineName,
		&discipline.Speciality,
		&discipline.Course)
	if err != nil {
		return nil, err
	}

	return discipline, nil
}
