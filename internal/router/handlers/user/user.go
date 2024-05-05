package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/router/handlers"
	"sbitnev_back/internal/router/handlers/encryption"
	"sbitnev_back/internal/router/middleware"
)

const (
	homePageUrl     = "/"
	LoginPageUrl    = "/login"
	journalPageUrl  = "/journal"
	schedulePageUrl = "/schedule"
	AdminMenuURL    = "/AdminMenu"
	StudentMenuURL  = "/StudentMenu"
	TeacherMenuURL  = "/TeacherMenu"
	ParentMenuURL   = "/ParentMenu"
	admin           = "admin"
	teacher         = "teacher"
	student         = "student"
	parent          = "parent"
)

func NewHandler(logger *slog.Logger, store *Store.Storage) handlers.Handler {
	return &handler{
		logger:  logger,
		storage: store,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET(homePageUrl, h.HomePage)
	router.GET(LoginPageUrl, h.LoginPage)
	router.POST(LoginPageUrl, h.UserIdent)
	//router.POST(homePageUrl, h.FeedBack)

	AdminMenuPath := router.Group("/adminPanel")
	AdminMenuPath.Use(middleware.CheckAdminAuth(h.storage))
	AdminMenuPath.GET("/menu")
	AdminMenuPath.GET("/management")
	AdminMenuPath.POST("/management/scheduleReg")
	AdminMenuPath.POST("/management/userReg")
	AdminMenuPath.PATCH("/management/gradesRef")
	AdminMenuPath.POST("/management/userDel")
	AdminMenuPath.POST("/management/bdBackUp")
	AdminMenuPath.GET("/journal")
	AdminMenuPath.GET("/schedule")

	TeacherMenuPath := router.Group("/teacherPanel")
	TeacherMenuPath.Use(middleware.LoginCheck(), middleware.RoleCheck())
	TeacherMenuPath.GET("/menu")
	TeacherMenuPath.GET("/journal")
	TeacherMenuPath.POST("/journal")
	TeacherMenuPath.GET("/schedule")

	StudentMenuPath := router.Group("/studentPanel")
	StudentMenuPath.Use(middleware.LoginCheck(), middleware.RoleCheck())
	StudentMenuPath.GET("/menu")
	StudentMenuPath.GET("/journal")
	StudentMenuPath.GET("/schedule")

	ParentMenuPath := router.Group("/parentPanel")
	ParentMenuPath.Use(middleware.LoginCheck(), middleware.RoleCheck())
	ParentMenuPath.GET("/menu")
	ParentMenuPath.GET("/journal")
	ParentMenuPath.GET("/schedule")
}

//Дописать в функцию Auth() обращение к БД + установку роли для пользователя
// + Redirect на страницу защищенной группы

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

//??? Обговорить/обдумать идею автоотправления формы на какой-то рабочий email
/*
func (h *handler) FeedBack(c *gin.Context) {
	FIO := c.PostForm("FIO")
	email := c.PostForm("email")
	msg := c.PostForm("msg")

	c.Status(http.StatusOK)
}
*/

//Не забыть про хедеры ...

func (h *handler) UserIdent(c *gin.Context) {
	login := c.PostForm("login")
	pass := c.PostForm("password")

	userRep := h.storage.User()
	userData, err := userRep.GetUserByLogin(login)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[UserIdent] error while identifing user: %s", err))
		switch {
		case errors.Is(err, sql.ErrNoRows):
			c.String(http.StatusUnauthorized, "Error: No such user", err)
			return
		default:
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}

	if pass != userData.Password {
		c.String(http.StatusForbidden, "error: wrong login or password")
		return
	}

	token, err := encryption.MakeToken(userData.UserId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.SetCookie("Authorization", token,
		86400, "/",
		"localhost", false, true)

	switch userData.Role {
	case admin:
		c.Redirect(http.StatusMovedPermanently, AdminMenuURL)
	case teacher:
		c.Redirect(http.StatusMovedPermanently, TeacherMenuURL)
	case student:
		c.Redirect(http.StatusMovedPermanently, StudentMenuURL)
	case parent:
		c.Redirect(http.StatusMovedPermanently, ParentMenuURL)
	}
}
