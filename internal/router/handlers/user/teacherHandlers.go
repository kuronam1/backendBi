package user

import (
	"fmt"
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
	c.HTML(http.StatusOK, "homepage_teacher.html", nil)
}

func (h *TeacherHandler) GetJournal(c *gin.Context) {
	const op = "TeacherHandlers.GetJournal"
	teacherID, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

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

	teacher, err := h.Storage.User().GetUserByID(teacherID.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	journal, err := h.Storage.Journal().GetGroupJournalByDiscipline(groupName, disciplineName)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	groupNames, err := h.Storage.Groups().GetAllTeachersGroups(teacherID.(int))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	disciplines, err := h.Storage.Disciplines().GetDisciplinesByTeacherId(teacherID.(int))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	lessonsTime, err := h.Storage.Schedule().GetAllGroupsLessonsOneDis(groupName, disciplineName)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.HTML(http.StatusOK, "journal_teacher.html", gin.H{
		"Journal":        journal,
		"LessonsTime":    lessonsTime,
		"Disciplines":    disciplines,
		"GroupsNames":    groupNames,
		"TeacherName":    teacher.FullName,
		"GroupName":      groupName,
		"UsedDiscipline": disciplineName,
		"Table":          true,
	})
}

func (h *TeacherHandler) GetPreJournal(c *gin.Context) {
	const op = "TeacherHandlers.GetPreJournal"
	teacherID, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	teacher, err := h.Storage.User().GetUserByID(teacherID.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	groupNames, err := h.Storage.Groups().GetAllTeachersGroups(teacherID.(int))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	disciplines, err := h.Storage.Disciplines().GetDisciplinesByTeacherId(teacherID.(int))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(200, "journal_teacher.html", gin.H{
		"GroupsNames": groupNames,
		"Disciplines": disciplines,
		"TeacherName": teacher.FullName,
		"Table":       false,
	})
}

func (h *TeacherHandler) AddGrade() gin.HandlerFunc {
	type request struct {
		StudentName    string `json:"studentName"`
		DisciplineName string `json:"disciplineID"`
		Level          string `json:"level"`
		Date           string `json:"date"`
		Comment        string `json:"comment,omitempty"`
	}
	const op = "TeacherHandlers.AddGrade"
	return func(c *gin.Context) {
		var req request
		if err := c.BindJSON(&req); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		h.Logger.Info("%v", req)

		user, err := h.Storage.User().GetUserByName(req.StudentName)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		date, err := time.Parse(time.DateOnly, req.Date)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		discipline, err := h.Storage.Disciplines().GetDisciplineByName(req.DisciplineName)

		grade := &models.Grade{
			StudentID:    user.UserID,
			DisciplineID: discipline.DisciplineID,
			Level:        req.Level,
			Date:         date,
			Comment:      req.Comment,
		}

		if err := h.Storage.Journal().CreateGrade(grade); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
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
	const op = "TeacherHandlers.GetSchedule"
	teacherID, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	teacher, err := h.Storage.User().GetUserByID(teacherID.(int))

	schedule, err := h.Storage.Schedule().GetScheduleByTeacherID(teacherID.(int))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.HTML(http.StatusOK, "schedule_teacher.html", gin.H{
		"Schedule": schedule,
		"FullName": teacher.FullName,
	})
}
