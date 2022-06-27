package app

import (
	"github.com/adityaeka26/golang-microservices/user/module/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler handler.Handler) *gin.Engine {
	router := gin.New()

	router.GET("/user", handler.GetUser)
	router.POST("/user", handler.CreateUser)

	return router
}
