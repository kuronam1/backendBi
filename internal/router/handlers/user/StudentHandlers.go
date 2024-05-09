package user

import (
	"log/slog"
	"sbitnev_back/internal/database/Store"
)

type StudentHandler struct {
	Logger  *slog.Logger
	Storage *Store.Storage
}
