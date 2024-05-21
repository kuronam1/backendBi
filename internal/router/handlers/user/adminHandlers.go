package user

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/database/models"
	"strconv"
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

	students, err := h.Storage.User().GetUserByRole("student")
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	teachers, err := h.Storage.User().GetAllTeachers()
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
		"Students":     students,
		"Groups":       groups,
		"Specialities": specialities,
		"Teachers":     teachers,
	})
}

// Registers

func (h *AdminHandler) ScheduleRegister(c *gin.Context) {
	const op = "AdminHandlers.ScheduleRegister"
	groupName := c.PostForm("group")
	if groupName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "no groupName",
		})
		return
	}

	h.Logger.Debug(groupName)

	if body, err := io.ReadAll(c.Request.Body); err != nil {
		panic(err)
	} else {
		c.Request.Body = io.NopCloser(bytes.NewReader(body))
	}

	form, _ := c.MultipartForm()
	for _, fh := range form.File["file"] {
		fileReader, _ := fh.Open()
		_, _ = io.ReadAll(fileReader)
		//log.Println("  file contents: ", string(contents))
		fileReader.Close()
	}

	file, err := c.FormFile("file")
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	h.Logger.Debug("file grabbed")

	filePath := fmt.Sprintf("./schedule/%v", file.Filename)

	h.Logger.Debug(fmt.Sprintf("filePath: %v", filePath))

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := h.Storage.Schedule().ScheduleRegister(filePath, groupName); err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "schedule registered",
	})
}

func (h *AdminHandler) UserRegister() gin.HandlerFunc {
	type Request struct {
		Login       string `json:"login"`
		Password    string `json:"password"`
		UserName    string `json:"userName"`
		Role        string `json:"role"`
		GroupName   string `json:"groupName,omitempty"`
		StudentName string `json:"studentName,omitempty"`
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

		switch user.Role {
		/*case "teacher":
		if err := rep.CreateTeacherDisciplineLink(id); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"StatusCode":  http.StatusInternalServerError,
				"Description": "Ошибка создания связи",
			})
			return
		}*/
		case "student":
			if err := rep.CreateGroupUserLink(id, request.GroupName); err != nil {
				h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
				c.HTML(http.StatusInternalServerError, "error.html", gin.H{
					"StatusCode":  http.StatusInternalServerError,
					"Description": "Ошибка создания связи",
				})
				return
			}
		case "parent":
			student, err := h.Storage.User().GetUserByName(request.StudentName)
			if err != nil {
				h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
				c.HTML(http.StatusInternalServerError, "error.html", gin.H{
					"StatusCode":  http.StatusInternalServerError,
					"Description": "Ошибка создания связи",
				})
				return
			}
			if err := rep.CreateParentStudentLink(id, student.UserID); err != nil {
				c.HTML(http.StatusInternalServerError, "error.html", gin.H{
					"StatusCode":  http.StatusInternalServerError,
					"Description": "Ошибка создания связи",
				})
				return
			}
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": "user registered",
		})
	}
}

func (h *AdminHandler) RecoverUserPassword() gin.HandlerFunc {
	type request struct {
		Login       string `json:"login"`
		NewPassword string `json:"newPassword"`
	}
	const op = "AdminHandler.RecoverUserPassword"
	return func(c *gin.Context) {
		var req request
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("cannot bind request [%s]", err),
			})
			return
		}

		if err := h.Storage.User().UpdateUserPassword(req.Login, req.NewPassword); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Status": "password successfully recovered",
		})
	}
}

func (h *AdminHandler) CreateDiscipline() gin.HandlerFunc {
	type request struct {
		TeacherName    string `json:"teacherName"`
		DisciplineName string `json:"disciplineName"`
		SpecialityName string `json:"specialityName"`
		Course         string `json:"course"`
	}
	const op = "AdminHandler.CreateDiscipline"
	return func(c *gin.Context) {
		var req request
		if err := c.BindJSON(&req); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err.Error()))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("cannot bind request [%s]", err.Error()),
			})
			return
		}

		h.Logger.Debug(fmt.Sprintf("%v", req))

		course, err := strconv.Atoi(req.Course)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err.Error()))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("cannot bind request [%s]", err.Error()),
			})
			return
		}

		if err = h.Storage.Disciplines().RegisterDiscipline(req.TeacherName, req.DisciplineName, req.SpecialityName, course); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err.Error()))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("cannot bind request [%s]", err.Error()),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": "discipline successfully registered",
		})
	}
}

func (h *AdminHandler) GroupRegister() gin.HandlerFunc {
	type request struct {
		Speciality string `json:"speciality"`
		Number     string `json:"number"`
		Course     string `json:"course"`
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

		h.Logger.Info("%v", req)

		number, err := strconv.Atoi(req.Number)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s in number", op, err))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		course, err := strconv.Atoi(req.Course)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s in course", op, err))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		group := &models.Group{
			Number:     number,
			Speciality: req.Speciality,
			Course:     course,
		}
		if err := h.Storage.Groups().GroupRegistration(group); err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": "group created",
		})
	}
}

func (h *AdminHandler) UpdateCourse(c *gin.Context) {
	const op = "AdminHandler.SwitchCourse"
	if err := h.Storage.Schedule().ClearSchedule(); err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if err := h.Storage.Journal().ClearJournal(); err != nil {
		h.Logger.Error(fmt.Sprintf("%s.ClearJournal - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if err := h.Storage.Groups().UpdateGroupsCourse(); err != nil {
		h.Logger.Error(fmt.Sprintf("%s.UpdateGroupsCourse - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	groups, err := h.Storage.Groups().GetAllGroups()
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s.GetAllGroups - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	for _, group := range groups {
		if group.Course >= 4 {
			if err := h.Storage.User().DeleteStudentsByGroup(group); err != nil {
				h.Logger.Error(fmt.Sprintf("%s.DeleteStudentsByGroup - %s", op, err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				return
			}
			if err := h.Storage.Groups().DeleteGroup(group); err != nil {
				h.Logger.Error(fmt.Sprintf("%s.DeleteGroup - %s", op, err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "course updated",
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
		"Journal":        journal,
		"GroupName":      groupName,
		"DisciplineName": disciplineName,
		"Lessons":        lessons,
		"Groups":         groups,
		"Disciplines":    disciplines,
		"Pre":            0,
	})
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
		"Pre":         1,
	})
	/*c.JSON(200, gin.H{
		"Groups":      groups,
		"Disciplines": disciplines,
	})*/
}

func (h *AdminHandler) GradesRefactor() gin.HandlerFunc {
	type Request struct {
		GradeID      string `json:"gradeID"`
		UserName     string `json:"name"`
		DisciplineID string `json:"discipline"`
		OldLevel     string `json:"oldLevel"`
		NewDate      string `json:"dateName"`
		NewLevel     string `json:"gradeName"`
		NewComment   string `json:"comment,omitempty"`
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

		h.Logger.Info(req.UserName)

		user, err := h.Storage.User().GetUserByName(req.UserName)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s in getUser", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		id, err := strconv.Atoi(req.DisciplineID)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s in getUser", op, err))
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

		Gradeid, err := strconv.Atoi(req.GradeID)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("%s - %s in parse", op, err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		oldGrade := &models.Grade{
			GradeID:      Gradeid,
			StudentID:    user.UserID,
			DisciplineID: id,
			Date:         newGradeDate,
			Level:        req.OldLevel,
		}

		NewGrade := &models.Grade{
			StudentID:    user.UserID,
			DisciplineID: id,
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
	groupName, existsG := c.GetQuery("group")
	teacherName, existsT := c.GetQuery("teacher")
	switch {
	case existsT && !existsG:
		h.ScheduleWithQueryTeacher(c, teacherName)
		return
	case !existsT && existsG:
		h.ScheduleWithQueryGroup(c, groupName)
		return
	default:
		h.GetPreSchedule(c)
		return
	}
}

func (h *AdminHandler) ScheduleWithQueryGroup(c *gin.Context, groupName string) {
	const op = "AdminHandlers.ScheduleWithQueryGroup"
	schedule, err := h.Storage.Schedule().GetScheduleByGroupName(groupName)
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
		"Schedule":  schedule,
		"Groups":    groups,
		"Teachers":  teachers,
		"Table":     "group",
		"GroupName": groupName,
	})
}

func (h *AdminHandler) ScheduleWithQueryTeacher(c *gin.Context, teacherName string) {
	const op = "AdminHandlers.ScheduleWithQueryTeacher"
	schedule, err := h.Storage.Schedule().GetScheduleByTeacherName(teacherName)
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
		"Schedule":    schedule,
		"Groups":      groups,
		"Teachers":    teachers,
		"Table":       "teacher",
		"TeacherName": teacherName,
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
		"Teachers":  teachers,
		"Groups":    groups,
		"ShowTable": "none",
	})
}
