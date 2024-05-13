package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"path/filepath"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/database/models"
	"time"
)

type AdminHandler struct {
	Logger  *slog.Logger
	Storage *Store.Storage
}

/*func (h *AdminHandler) Menu(c *gin.Context) {
	c.HTML(200, "", nil)
}*/

//"homePage"

func (h *AdminHandler) Management(c *gin.Context) {
	const op = "AdminHandlers.Management"
	groups, err := h.Storage.Groups().GetAllGroups()
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	specialities, err := h.Storage.Groups().GetAllSpecialities()
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "admin_management.html", gin.H{
		"Groups":       groups,
		"Specialities": specialities,
	})
}

// Registers

func (h *AdminHandler) ScheduleRegister(c *gin.Context) {
	const op = "AdminHandlers.ScheduleRegister"
	file, err := c.FormFile("file")
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	filePath := filepath.Join("/schedule", file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := h.Storage.Schedule().ScheduleRegister(filePath); err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(202, gin.H{
		"status": "schedule registered",
	})
}

func (h *AdminHandler) UserRegister() gin.HandlerFunc {
	type Request struct {
		Login     string `json:"login"`
		Password  string `json:"password"`
		UserName  string `json:"userName"`
		Role      string `json:"role"`
		GroupName string `json:"groupName,omitempty"`
	}
	return func(c *gin.Context) {
		const op = "AdminHandlers.UserRegister"
		rep := h.Storage.User()

		var request Request
		if err := c.BindJSON(&request); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.AbortWithStatusJSON(500, gin.H{
				"error": err,
			})
			return
		}

		user := &models.User{
			Login:    request.Login,
			Password: request.Password,
			FullName: request.UserName,
			Role:     request.Role,
		}
		id, err := rep.CreateUser(user)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.AbortWithStatusJSON(500, gin.H{
				"error": err,
			})
			return
		}

		if err := rep.CreateUserLink(id, request.GroupName); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"status": "user registered",
		})
	}
}

func (h *AdminHandler) GroupRegister() gin.HandlerFunc {
	type request struct {
		Speciality string `json:"speciality"`
		Number     int    `json:"number"`
		Course     int    `json:"course"`
	}
	return func(c *gin.Context) {
		const op = "AdminHandlers.GroupRegister"
		var req request
		if err := c.BindJSON(&req); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		group := &models.Group{
			Number:     req.Number,
			Speciality: req.Speciality,
			Course:     req.Course,
		}
		if err := h.Storage.Groups().GroupRegistration(group); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "group created",
		})
	}
}

func (h *AdminHandler) BackUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "done",
	})
}

// Grades and Journal

func (h *AdminHandler) GetJournal(c *gin.Context) {
	const op = "AdminHandlers.GetJournal"
	groupName, exist := c.GetQuery("group")
	if !exist {
		h.GetPreJournal(c)
		c.Abort()
		return
	}

	disciplineName, exist := c.GetQuery("discipline")
	if !exist {
		h.GetPreJournal(c)
		c.Abort()
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

	lessons, err := h.Storage.Schedule().GetAllGroupsLessonsOneDis(groupName, disciplineName)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	groups, err := h.Storage.Groups().GetAllGroups()
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
		return
	}

	disciplines, err := h.Storage.Disciplines().GetAllDisciplines()
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "admin_journal.html", gin.H{
		"Journal":     journal,
		"Lessons":     lessons,
		"Groups":      groups,
		"Disciplines": disciplines,
	})
	/*c.JSON(200, gin.H{
		"Journal": journal,
		"Lessons": lessons,
	})*/
}

func (h *AdminHandler) GetPreJournal(c *gin.Context) {
	const op = "AdminHandlers.GetPreJournal"
	groups, err := h.Storage.Groups().GetAllGroups()
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
		return
	}

	disciplines, err := h.Storage.Disciplines().GetAllDisciplines()
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "admin_journal.html", gin.H{
		"Groups":      groups,
		"Disciplines": disciplines,
	})
	/*c.JSON(200, gin.H{
		"Groups":      groups,
		"Disciplines": disciplines,
	})*/
}

func (h *AdminHandler) GradesRefactor() gin.HandlerFunc {
	type Request struct {
		GradeID        int    `json:"gradeID"`
		UserName       string `json:"userName"`
		DisciplineName string `json:"disciplineName"`
		OldLevel       string `json:"oldLevel"`
		OldDate        string `json:"oldDate"`
		OldComment     string `json:"oldComment,omitempty"`
		NewLevel       string `json:"newLevel"`
		NewDate        string `json:"newDate"`
		NewComment     string `json:"newComment,omitempty"`
	}
	return func(c *gin.Context) {
		const op = "AdminHandlers.GradesRefactor"
		var req Request
		if err := c.BindJSON(&req); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		user, err := h.Storage.User().GetUserByName(req.UserName)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s in getUser", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		discipline, err := h.Storage.Disciplines().GetDisciplineByName(req.DisciplineName)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s getDis", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		oldGradeDate, err := time.Parse(time.DateOnly, req.OldDate)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s in parse", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		newGradeDate, err := time.Parse(time.DateOnly, req.NewDate)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s in parse", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		oldGrade := &models.Grade{
			GradeID:      req.GradeID,
			StudentID:    user.UserID,
			DisciplineID: discipline.DisciplineID,
			Level:        req.OldLevel,
			Date:         oldGradeDate,
			Comment:      req.OldComment,
		}

		NewGrade := &models.Grade{
			StudentID:    user.UserID,
			DisciplineID: discipline.DisciplineID,
			Level:        req.NewLevel,
			Date:         newGradeDate,
			Comment:      req.NewComment,
		}

		if err := h.Storage.Journal().UpdateGrade(oldGrade, NewGrade); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s in ref", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"Accepted": "Grade is refactored",
		})
	}
}

// Schedules

func (h *AdminHandler) GetSchedule(c *gin.Context) {
	const op = "AdminHandlers.GetSchedule"
	groupName, exists := c.GetQuery("group")
	if exists && groupName != "" {
		h.ScheduleWithQueryGroup(c, groupName)
		c.Abort()
		return
	}

	teacherName, exists := c.GetQuery("teacher")
	if exists && teacherName != "" {
		h.ScheduleWithQueryTeacher(c, teacherName)
		c.Abort()
		return
	}

	h.GetPreSchedule(c)
}

func (h *AdminHandler) ScheduleWithQueryGroup(c *gin.Context, groupName string) {
	const op = "AdminHandlers.ScheduleWithQueryGroup"
	schedule, err := h.Storage.ScheduleMethods.GetScheduleByGroupName(groupName)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	case err != nil:
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	groups, err := h.Storage.Groups().GetAllGroups()
	switch {
	case errors.Is(err, sql.ErrNoRows):
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	case err != nil:
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	teachers, err := h.Storage.User().GetAllTeachers()
	switch {
	case errors.Is(err, sql.ErrNoRows):
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	case err != nil:
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "admin_schedule.html", gin.H{
		"Schedule": schedule,
		"Groups":   groups,
		"Teachers": teachers,
	})
}

func (h *AdminHandler) ScheduleWithQueryTeacher(c *gin.Context, teacherName string) {
	const op = "AdminHandlers.ScheduleWithQueryTeacher"
	schedule, err := h.Storage.ScheduleMethods.GetScheduleByTeacherName(teacherName)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	case err != nil:
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	groups, err := h.Storage.Groups().GetAllGroups()
	switch {
	case errors.Is(err, sql.ErrNoRows):
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	case err != nil:
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	teachers, err := h.Storage.User().GetAllTeachers()
	switch {
	case errors.Is(err, sql.ErrNoRows):
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	case err != nil:
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "admin_schedule.html", gin.H{
		"Schedule": schedule,
		"Groups":   groups,
		"Teachers": teachers,
	})
}

func (h *AdminHandler) GetPreSchedule(c *gin.Context) {
	const op = "AdminHandlers.GetPreSchedule"
	groups, err := h.Storage.Groups().GetAllGroups()
	switch {
	case errors.Is(err, sql.ErrNoRows):
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	case err != nil:
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	teachers, err := h.Storage.User().GetAllTeachers()
	switch {
	case errors.Is(err, sql.ErrNoRows):
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	case err != nil:
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(200, "admin_schedule.html", gin.H{
		"Teachers": teachers,
		"Groups":   groups,
	})
}
