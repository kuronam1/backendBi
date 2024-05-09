package Store

import (
	"database/sql"
	"errors"
	"sbitnev_back/internal/database/models"
)

type DisciplineRepository struct {
	store *Storage
}

func (d *DisciplineRepository) GetDisciplineByName(name string) (*models.Discipline, error) {
	const op = "fc.journalRep.GetDisciplineByName"
	stmt, err := d.store.DB.Prepare("SELECT teacher_id, discipline_name, speciality, course FROM disciplines WHERE discipline_name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var dis = models.Discipline{
		DisciplineName: name,
	}
	err = stmt.QueryRow(name).Scan(
		&dis.DisciplineID,
		&dis.TeacherID,
		&dis.DisciplineName,
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

func (d *DisciplineRepository) RegisterDiscipline(teacherName, disciplineName, speciality string, course int) (int64, error) {
	const op = "fc.discRep.RegisterDiscipline"
	stmt, err := d.store.DB.Prepare("INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING;")
	if err != nil {
		return 0, err
	}

	teacher, err := d.store.User().GetUserByName(teacherName)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(teacher.UserID, disciplineName, speciality, course)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	return id, nil
}