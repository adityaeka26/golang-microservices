package router

import (
	"github.com/adityaeka26/golang-microservices/notification/module/handler"
	"github.com/gin-gonic/gin"
)

type Router interface {
	GetGinEngine() *gin.Engine
}

type RouterImpl struct {
	ginEngine *gin.Engine
}

func NewRouter(handler handler.Handler) Router {
	router := gin.New()

	go handler.SendRegisterOtp()

	return &RouterImpl{
		ginEngine: router,
	}
}

func (router RouterImpl) GetGinEngine() *gin.Engine {
	return router.ginEngine
}
