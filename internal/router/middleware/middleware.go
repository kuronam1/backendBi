package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/router/handlers/encryption"
)

var (
	notValidLoginPass = errors.New("not valid data")
	notValidRole      = errors.New("role is not valid")
)

const (
	Admin   = "admin"
	Teacher = "teacher"
	Student = "student"
	Parent  = "parent"
)

func CheckAdminAuth(storage *Store.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "not authorized",
			})
			return
		} else if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "not authorized",
			})
			return
		}

		id, err := encryption.ParsingToken(token)
		switch {
		case errors.Is(err, encryption.NotValid):
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "not valid token",
			})
			return
		case errors.Is(err, encryption.ParsingErr):
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
		userRep := storage.User()
		user, err := userRep.GetUserByID(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err,
			})
			return
		}

		if user.Role != Admin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid role",
			})
			return
		}

		c.Next()
	}
}

func CheckStudentAuth(storage *Store.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "not authorized",
			})
			return
		} else if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "not authorized",
			})
			return
		}

		id, err := encryption.ParsingToken(token)
		switch {
		case errors.Is(err, encryption.NotValid):
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "not valid token",
			})
			return
		case errors.Is(err, encryption.ParsingErr):
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
		userRep := storage.User()
		user, err := userRep.GetUserByID(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err,
			})
			return
		}

		if user.Role != Student {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid role",
			})
			return
		}
		c.Set("id", id)

		c.Next()
	}
}

func CheckTeacherAuth(storage *Store.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "not authorized",
			})
			return
		} else if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "not authorized",
			})
			return
		}

		id, err := encryption.ParsingToken(token)
		switch {
		case errors.Is(err, encryption.NotValid):
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "not valid token",
			})
			return
		case errors.Is(err, encryption.ParsingErr):
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		user, err := storage.User().GetUserByID(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err,
			})
			return
		}

		if user.Role != Teacher {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid role",
			})
			return
		}
		c.Set("id", id)

		c.Next()
	}
}

func CheckParentAuth(storage *Store.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "not authorized",
			})
			return
		} else if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "not authorized",
			})
			return
		}

		id, err := encryption.ParsingToken(token)
		switch {
		case errors.Is(err, encryption.NotValid):
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "not valid token",
			})
			return
		case errors.Is(err, encryption.ParsingErr):
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		user, err := storage.User().GetUserByID(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err,
			})
			return
		}

		if user.Role != Parent {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid role",
			})
			return
		}
		c.Set("id", id)

		c.Next()
	}
}

func LoggingReq(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info(fmt.Sprintf("[request] Method: %v, addr: %v", c.Request.Method, c.Request.URL))
		c.Next()
	}
}
