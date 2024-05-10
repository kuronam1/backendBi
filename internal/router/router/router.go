package router

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/router/handlers/user"
	"sbitnev_back/internal/router/middleware"
)

func InitRouter(log *slog.Logger, db *Store.Storage) http.Handler {
	log.Info("[router]: registration front files and routes")
	router := gin.New()
	router.Use(gin.Recovery(), middleware.LoggingReq(log))
	/*router.LoadHTMLGlob("html/*")
	router.Static("/static/assets", "./static/assets")*/

	handler := user.NewHandler(log, db)
	handler.Register(router)
	return router
}
