package router

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"path/filepath"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/router/handlers/user"
	"sbitnev_back/internal/router/middleware"
)

func InitRouter(log *slog.Logger, db *Store.Storage) http.Handler {
	log.Info("[router]: registration front files and routes")
	router := gin.New()
	router.Use(gin.Recovery(), middleware.LoggingReq(log))
	router.LoadHTMLGlob(filepath.Join("static", "html", "homePage.html"))

	handler := user.NewHandler(log, db)
	handler.Register(router)
	return router
}
