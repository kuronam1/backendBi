package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/database/Store"
)

type ParentHandler struct {
	Logger  *slog.Logger
	Storage *Store.Storage
}

func (h *ParentHandler) Menu(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage_parent.html", nil)
}

func (h *ParentHandler) GetSchedule(c *gin.Context) {
	const op = "ParentHandler.GetSchedule"
	parentId, exists := c.Get("id")
	if !exists {
		h.Logger.Error(fmt.Sprintf("%s - cannot get parentID", op))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	studentId, err := h.Storage.User().GetStudentIDByParentID(parentId.(int))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	group, err := h.Storage.Groups().GroupMembership(studentId)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	schedule, err := h.Storage.Schedule().GetScheduleByGroupName(group.Name)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "schedule_parent.html", gin.H{
		"Schedule":  schedule,
		"GroupName": group.Name,
	})
}

func (h *ParentHandler) GetJournal(c *gin.Context) {
	const op = "StudentHandlers.GetJournal"
	parentId, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	studentId, err := h.Storage.User().GetStudentIDByParentID(parentId.(int))

	user, err := h.Storage.User().GetUserByID(studentId)

	journal, err := h.Storage.Journal().GetJournalByStudentID(studentId)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "journal_parent.html", gin.H{
		"Journal":     journal,
		"StudentName": user.FullName,
	})
}
