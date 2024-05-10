package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/router/handlers/encryption"
)

var (
	notValidLoginPass = errors.New("not valid data")
	notValidRole      = errors.New("role is not valid")
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

		if user.Role != "admin" {
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

		if user.Role != "student" {
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

		if user.Role != "teacher" {
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

		if user.Role != "parent" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid role",
			})
			return
		}
		c.Set("id", id)

		c.Next()
	}
}
