package logging

import (
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

func InitLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case EnvLocal:
		logger = slog.New(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
	case EnvProd:
		logger = slog.New(slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
