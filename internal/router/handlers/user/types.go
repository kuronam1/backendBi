package user

import "log/slog"

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type handler struct {
	logger *slog.Logger
}
