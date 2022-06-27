package router

import "github.com/gin-gonic/gin"

type Router interface {
	GetGinEngine() *gin.Engine
}
