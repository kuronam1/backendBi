package user

import (
	"log/slog"
	"sbitnev_back/internal/database/Store"
)

type Userdata struct {
	Login    string
	Password string
}

type handler struct {
	logger *slog.Logger
	db     *Store.Storage
}
