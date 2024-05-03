package scheduleRepository

import (
	"database/sql"
)

type ScheduleRepository struct {
	DB *sql.DB
}
