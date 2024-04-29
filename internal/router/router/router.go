package router

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"path/filepath"
	"sbitnev_back/internal/router/handlers/user"
	"sbitnev_back/internal/router/middleware"
)

func InitRouter(log *slog.Logger) http.Handler {
	log.Info("[router]: registration front files and routes")
	router := gin.New()
	router.Use(gin.Recovery(), middleware.LoggingReq(log)) // change logger!!
	router.LoadHTMLGlob(filepath.Join("static", "html", "homePage.html"))

	handler := user.NewHandler(log)
	handler.Register(router)
	return router
}
