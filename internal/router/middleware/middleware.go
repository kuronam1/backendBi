package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sbitnev_back/internal/database/Store"
	"sbitnev_back/internal/router/handlers/encryption"
	"slices"
)

var (
	notValidLoginPass = errors.New("not valid data")
	notValidRole      = errors.New("role is not valid")
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.Abort()
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			//указать ошибку! token not valid!
			c.Abort()
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		//Осуществить подгрузку с БД токена
		//для его последующей проверка
		//Оставляю затычку в виде несуществующего токена
		token := authHeader[7:]
		if token != "Token" {
			//указать ошибку! token not valid!
			c.Abort()
			c.Redirect(http.StatusMovedPermanently, "/login")
		}

		c.Next()
	}
}

//Установить хедеры,описать ошибку в func RoleCheck()

func RoleCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleHeader := c.GetHeader("Role")
		err := redirectFunc(roleHeader)
		if err != nil {
			//c.AbortWithError(http.StatusForbidden, err)
			c.Redirect(http.StatusFound, "/login")
			return
		}

		c.Next()
	}
}

// Исправить костыль с массивом ролей!
func redirectFunc(header string) error {
	if header == "" {
		return notValidRole
	}

	roles := []string{"admin", "student", "teacher"}

	if !slices.Contains(roles, header) {
		return notValidRole
	}
	return nil
}

func LoggingReq(l *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		l.Info(fmt.Sprintf("recieved a req: Method - %s, Addr - %s", c.Request.Method, c.Request.URL.Path))
	}
}

func CheckAdminAuth(storage *Store.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authorization")
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		} else if token == "" {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}

		id, err := encryption.ParsingToken(token)
		userRep := storage.User()
		user, err := userRep.GetUserByID(id)
		switch {
		case errors.Is(err, encryption.NotValid):
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "not valid token",
			})
		case errors.Is(err, encryption.ParsingErr):
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid role",
			})
		}

		c.Next()
	}
}

//Действие - загрузка расписания (метод) (что надо с фронта)
// + сделать функции под бд
