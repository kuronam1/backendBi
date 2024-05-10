package user

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/database/models"
	"time"
)

type TeacherHandler struct {
	Logger  *slog.Logger
	Storage *Store.Storage
}

func (h *TeacherHandler) Menu(c *gin.Context) {
	c.HTML(http.StatusOK, "", nil)
}

func (h *TeacherHandler) GetJournal(c *gin.Context) {
	groupName, exist := c.GetQuery("group")
	if !exist {
		h.GetPreJournal(c)
		return
	}

	disciplineName, exist := c.GetQuery("discipline")
	if !exist {
		h.GetPreJournal(c)
		return
	}

	journal, err := h.Storage.Journal().GetGroupJournalByDiscipline(groupName, disciplineName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "", gin.H{
		"journal": journal,
	})
}

func (h *TeacherHandler) GetPreJournal(c *gin.Context) {
	groups, err := h.Storage.Groups().GetAllGroups()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
		return
	}

	disciplines, err := h.Storage.Disciplines().GetAllDisciplines()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(200, "", gin.H{
		"groups":      groups,
		"disciplines": disciplines,
	})
}

//Узнать и исправить,если будет баг с парсом времени

func (h *TeacherHandler) AddGrade() gin.HandlerFunc {
	type request struct {
		StudentName    string `json:"studentName"`
		DisciplineName string `json:"disciplineName"`
		Level          int    `json:"level"`
		Date           string `json:"date"`
		Comment        string `json:"comment,omitempty"`
	}
	return func(c *gin.Context) {
		var req request
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		user, err := h.Storage.User().GetUserByName(req.StudentName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		discipline, err := h.Storage.Disciplines().GetDisciplineByName(req.DisciplineName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		date, err := time.Parse(time.DateOnly, req.Date)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		grade := &models.Grade{
			StudentID:    user.UserID,
			DisciplineID: discipline.DisciplineID,
			Level:        req.Level,
			Date:         date,
			Comment:      req.Comment,
		}

		if err := h.Storage.Journal().CreateGrade(grade); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "grade is created",
		})
	}
}

func (h *TeacherHandler) GetSchedule(c *gin.Context) {
	teacherID, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	user, err := h.Storage.User().GetUserByID(teacherID.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	schedule, err := h.Storage.Schedule().GetScheduleByTeacherName(user.FullName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "", gin.H{
		"Schedule": schedule,
	})
}
