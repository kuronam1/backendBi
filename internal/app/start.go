package app

import (
	"fmt"
	"os"
	"os/signal"
	"sbitnev_back/errors"
	"sbitnev_back/internal/config"
	rt "sbitnev_back/internal/router/router"
	"sbitnev_back/pkg/logging"
	"sbitnev_back/pkg/server"
	"syscall"
)

func Run(cfg *config.Config) {
	logger := logging.InitLogger(cfg.Env)

	router := rt.InitRouter(logger)

	logger.Info(fmt.Sprintf("[Server] Starting server on port %s", cfg.Addr))
	httpServer := server.New(router, cfg)
	logger.Info(fmt.Sprintf("[Server] Server started on port %s", cfg.Addr))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		logger.Error(fmt.Sprintf("[Server] run signal: %s", s.String()))
	case err := <-httpServer.Notify:
		logger.Error(fmt.Sprintf(
			"[Server] run http.Server.Notify: %s",
			errors.Error(err)))
	}

	logger.Error(fmt.Sprintf("[Server]: Shutting down http.Server in %.1f", cfg.ShutDownTimout.Seconds()))
	err := httpServer.Shutdown()
	if err != nil {
		logger.Error(fmt.Sprintf("[Server] Stopped - http.Server.Shutdown: %s", errors.Error(err)))
	}
}
