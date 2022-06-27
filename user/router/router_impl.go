package router

import (
	"github.com/adityaeka26/golang-microservices/user/module/handler"
	"github.com/gin-gonic/gin"
)

type RouterImpl struct {
	ginEngine *gin.Engine
}

func NewRouter(handler handler.Handler) Router {
	router := gin.New()

	router.GET("/user", handler.GetUser)
	router.POST("/user", handler.CreateUser)

	return &RouterImpl{
		ginEngine: router,
	}
}

func (router RouterImpl) GetGinEngine() *gin.Engine {
	return router.ginEngine
}
