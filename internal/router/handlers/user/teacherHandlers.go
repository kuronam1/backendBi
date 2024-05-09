package user

import (
	"log/slog"
	"sbitnev_back/internal/database/Store"
)

type TeacherHandler struct {
	Logger  *slog.Logger
	Storage *Store.Storage
}
