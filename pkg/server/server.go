package server

import (
	"context"
	"net/http"
	"sbitnev_back/internal/config"
	"time"
)

type App struct {
	Server       *http.Server
	Notify       chan error
	shutdownTime time.Duration
}

func New(handler http.Handler, cfg *config.Config) *App {

	s := &http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}

	app := &App{
		Server:       s,
		Notify:       make(chan error),
		shutdownTime: cfg.ShutDownTimout,
	}

	app.Start()

	return app
}

func (app *App) Start() {
	go func() {
		app.Notify <- app.Server.ListenAndServe()
		close(app.Notify)
	}()
}

func (app *App) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), app.shutdownTime)
	defer cancel()

	return app.Server.Shutdown(ctx)
}
