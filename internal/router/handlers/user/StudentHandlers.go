package user

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/database/Store"
)

type StudentHandler struct {
	Logger  *slog.Logger
	Storage *Store.Storage
}

func (h *StudentHandler) Menu(c *gin.Context) {
	c.HTML(http.StatusOK, "", nil)
}

func (h *StudentHandler) GetSchedule(c *gin.Context) {
	studentId, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	schedule, err := h.Storage.Schedule().GetScheduleByID(studentId.(int))
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

func (h *StudentHandler) GetJournal(c *gin.Context) {
	studentId, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	journal, err := h.Storage.Journal().GetJournalByStudentID(studentId.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.HTML(http.StatusOK, "", gin.H{
		"Journal": journal,
	})
}
