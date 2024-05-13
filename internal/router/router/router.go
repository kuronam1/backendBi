package router

import (
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/gin-gonic/gin"
	"html/template"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/database/models"
	"sbitnev_back/internal/router/handlers/user"
	"sbitnev_back/internal/router/middleware"
)

func InitRouter(log *slog.Logger, db *Store.Storage) http.Handler {
	log.Info("[router]: registration front files and routes")
	router := gin.New()
	router.Use(gin.Recovery(), middleware.LoggingReq(log))
	//router.LoadHTMLGlob("html/*")
	router.SetFuncMap(template.FuncMap{
		"Get": GetFunc,
	})
	router.LoadHTMLFiles("html/admin_journal.html", "html/admin_schedule.html", "html/index.html", "html/admin_management.html")
	router.Static("/static", "./static")

	handler := user.NewHandler(log, db)
	handler.Register(router)
	return router
}

func GetFunc(key string, m *treemap.Map) []models.Grade {
	res, exist := m.Get(key)
	if !exist {
		return nil
	}
	return res.([]models.Grade)
}
