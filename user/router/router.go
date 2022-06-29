package router

import (
	"github.com/adityaeka26/golang-microservices/user/middleware"
	"github.com/adityaeka26/golang-microservices/user/module/handler"
	"github.com/gin-gonic/gin"
)

type Router interface {
	GetGinEngine() *gin.Engine
}

type RouterImpl struct {
	ginEngine *gin.Engine
	auth      middleware.Auth
}

func NewRouter(handler handler.Handler, auth middleware.Auth) Router {
	router := gin.New()

	userRouter := router.Group("/user")

	userRouter.POST("/v1/register", handler.Register)
	userRouter.POST("/v1/register/verify", handler.VerifyRegister)

	return &RouterImpl{
		ginEngine: router,
	}
}

func (router RouterImpl) GetGinEngine() *gin.Engine {
	return router.ginEngine
}
