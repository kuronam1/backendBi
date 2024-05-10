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
	homePageUrl    = "/"
	LoginPageUrl   = "/login"
	LogoutPageURl  = "/logout"
	AdminMenuURL   = "/AdminMenu"
	StudentMenuURL = "/StudentMenu"
	TeacherMenuURL = "/TeacherMenu"
	ParentMenuURL  = "/ParentMenu"
	admin          = "admin"
	teacher        = "teacher"
	student        = "student"
	parent         = "parent"
)

func NewHandler(logger *slog.Logger, store *Store.Storage) handlers.Handler {
	return &handler{
		logger:  logger,
		storage: store,
		AdminHandler: &AdminHandler{
			Logger:  logger,
			Storage: store,
		},
		StudentHandler: &StudentHandler{
			Logger:  logger,
			Storage: store,
		},
		TeacherHandler: &TeacherHandler{
			Logger:  logger,
			Storage: store,
		},
	}
}

//Дописать проверки middleware !!!!

func (h *handler) Register(router *gin.Engine) {
	router.GET(homePageUrl, h.HomePage)
	router.POST(LoginPageUrl, h.UserIdent())
	router.POST(LogoutPageURl, h.Logout)
	//router.POST(homePageUrl, h.FeedBack)

	AdminMenuPath := router.Group("/adminPanel")
	AdminMenuPath.Use(middleware.CheckAdminAuth(h.storage))
	//AdminMenuPath.GET("/menu", h.AdminHandler.Menu)
	AdminMenuPath.GET("/management", h.AdminHandler.Management)
	AdminMenuPath.POST("/management/scheduleReg", h.AdminHandler.ScheduleRegister)
	AdminMenuPath.POST("/management/userReg", h.AdminHandler.UserRegister())
	//AdminMenuPath.POST("/management/userDel")
	//AdminMenuPath.PATCH("/management/userRef")
	AdminMenuPath.POST("/management/groupReg", h.AdminHandler.GroupRegister())
	AdminMenuPath.PATCH("/management/gradesRef", h.AdminHandler.GradesRefactor())
	AdminMenuPath.POST("/management/bdBackUp", h.AdminHandler.BackUp)
	AdminMenuPath.GET("/journal", h.AdminHandler.GetJournal)
	AdminMenuPath.GET("/schedule", h.AdminHandler.GetSchedule)

	TeacherMenuPath := router.Group("/teacherPanel")
	TeacherMenuPath.Use(middleware.CheckTeacherAuth(h.storage))
	TeacherMenuPath.GET("/menu", h.TeacherHandler.Menu)
	TeacherMenuPath.GET("/journal", h.TeacherHandler.GetJournal)
	TeacherMenuPath.POST("/journal", h.TeacherHandler.AddGrade())
	TeacherMenuPath.GET("/schedule", h.TeacherHandler.GetSchedule)

	StudentMenuPath := router.Group("/studentPanel")
	StudentMenuPath.Use(middleware.CheckStudentAuth(h.storage))
	StudentMenuPath.GET("/menu", h.StudentHandler.Menu)
	StudentMenuPath.GET("/journal", h.StudentHandler.GetJournal)
	StudentMenuPath.GET("/schedule", h.StudentHandler.GetSchedule)

	ParentMenuPath := router.Group("/parentPanel")
	ParentMenuPath.Use(middleware.CheckParentAuth(h.storage))
	ParentMenuPath.GET("/menu", h.StudentHandler.Menu)
	ParentMenuPath.GET("/journal", h.StudentHandler.GetJournal)
	ParentMenuPath.GET("/schedule", h.StudentHandler.GetSchedule)
}

func (h *handler) HomePage(c *gin.Context) {
	if c.FullPath() != homePageUrl {
		c.HTML(404, "", nil)
	}
	c.HTML(200, "index.tmpl", nil)
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

func (h *handler) UserIdent() gin.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	return func(c *gin.Context) {
		var req request
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err in bind": err,
			})
			return
		}

		userRep := h.storage.User()
		userData, err := userRep.GetUserByLogin(req.Login)
		if err != nil {
			h.logger.Error(fmt.Sprintf("[UserIdent] error while identifing user: %s", err))
			switch {
			case errors.Is(err, sql.ErrNoRows):
				c.String(http.StatusUnauthorized, "Error: No such user", err)
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"err in getUser": err,
				})
				return
			}
		}

		if req.Password != userData.Password {
			c.String(http.StatusForbidden, "error: wrong login or password")
			return
		}

		token, err := encryption.MakeToken(userData.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err in makeToken": err,
			})
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
}

func (h *handler) Logout(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	_, err = encryption.ParsingToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": encryption.NotValid,
		})
	}

	c.SetCookie("Authorization", token,
		-1, "/",
		"localhost", false, true)

	c.Redirect(http.StatusMovedPermanently, homePageUrl)
}
