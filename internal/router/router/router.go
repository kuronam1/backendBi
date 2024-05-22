package router

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/database/models"
	"sbitnev_back/internal/router/handlers/user"
	"sbitnev_back/internal/router/middleware"
	"strconv"
	"time"
)

const timePayload = "02-01-2006"

func InitRouter(log *slog.Logger, db *Store.Storage) http.Handler {
	log.Info("[router]: registration front files and routes")
	router := gin.New()
	router.Use(gin.Recovery(), middleware.LoggingReq(log))
	//router.LoadHTMLGlob("html/*")
	router.SetFuncMap(template.FuncMap{
		"TFormat": TFormat,
		"Count":   CountN,
		"Avg":     Avg,
	})
	router.LoadHTMLGlob("html/*")
	router.Static("/static", "./static")

	handler := user.NewHandler(log, db)
	handler.Register(router)
	return router
}

func TFormat(t time.Time) string {
	return t.Format(timePayload)
}

func CountN(grades []models.Grade) int {
	counter := 0
	for _, value := range grades {
		if value.Level == "н" {
			counter++
		}
	}
	return counter
}

func Avg(grades []models.Grade) float64 {
	sum := 0
	for _, value := range grades {
		if value.Level != "н" {
			num, _ := strconv.Atoi(value.Level)
			sum += num
		}
	}
	return float64(sum) / float64(len(grades))
}
