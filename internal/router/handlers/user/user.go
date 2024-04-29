package user

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/router/handlers"
	"sbitnev_back/internal/router/middleware"
)

const (
	homePageUrl     = "/"
	LoginPageUrl    = "/login"
	journalPageUrl  = "/journal"
	schedulePageUrl = "/schedule"
	token           = "Token"
	LoginPassString = "MyLoginmypassword"
)

func NewHandler(logger *slog.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(homePageUrl, h.HomePage)
	router.GET(LoginPageUrl, h.LoginPage)
	router.POST(LoginPageUrl, h.Auth)

	protectedPath := router.Group(LoginPageUrl)
	protectedPath.Use(middleware.LoginCheck(), middleware.RoleCheck())
	protectedPath.GET(journalPageUrl, h.JournalPage)
	protectedPath.GET(schedulePageUrl, h.SchedulePage)
}

//Дописать в функцию Auth() обращение к БД + установку роли для пользователя
// + Redirect на страницу защищенной группы

func (h *handler) Auth(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	if login == "" || password == "" {
		c.String(http.StatusNotAcceptable, "wrong login or password")
		return
	}

	if login+password != LoginPassString {
		c.String(http.StatusNotAcceptable, "invalid login or password")
		return
	}

	c.Header("Authorization", token)
	c.String(http.StatusOK, "auth passed")
}

func (h *handler) HomePage(c *gin.Context) {
	if c.FullPath() != homePageUrl {
		c.HTML(404, "", nil)
	}
	c.HTML(200, "homePage.html", nil)
}

func (h *handler) LoginPage(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"status": "OK! u'r on auth page",
	})
	c.Writer.Pusher()
}

func (h *handler) SchedulePage(c *gin.Context) {
	c.JSON(http.StatusAccepted, map[string]string{
		"status":       "Accepted! u'r on schedule page",
		"your role is": c.GetHeader("role"),
	})
}

func (h *handler) JournalPage(c *gin.Context) {
	c.JSON(http.StatusAccepted, map[string]string{
		"status": "Accepted! u'r on journal page",
	})
}
