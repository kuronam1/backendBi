package handlers

import "github.com/gin-gonic/gin"

type Router interface {
	initRouter() *gin.Engine
}
