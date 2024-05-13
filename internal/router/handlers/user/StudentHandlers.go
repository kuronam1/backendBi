package user

import (
	"fmt"
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
	//c.HTML(http.StatusOK, "", nil)
	c.JSON(http.StatusOK, gin.H{
		"OK": "u r in menu",
	})
}

func (h *StudentHandler) GetSchedule(c *gin.Context) {
	const op = "StudentHandlers.GetSchedule"
	studentId, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	group, err := h.Storage.Groups().GroupMembership(studentId.(int))
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

	/*c.HTML(http.StatusOK, "", gin.H{
		"Schedule": schedule,
	})*/
	c.JSON(http.StatusOK, gin.H{
		"Schedule": schedule,
	})
}

func (h *StudentHandler) GetJournal(c *gin.Context) {
	const op = "StudentHandlers.GetJournal"
	studentId, exists := c.Get("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot identify",
		})
		return
	}

	journal, err := h.Storage.Journal().GetJournalByStudentID(studentId.(int))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("%s - %s", op, err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	/*c.HTML(http.StatusOK, "", gin.H{
		"Journal": journal,
	})*/
	c.JSON(http.StatusOK, gin.H{
		"Journal": journal,
	})
}
