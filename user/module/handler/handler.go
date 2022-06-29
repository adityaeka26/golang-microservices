package handler

import (
	"net/http"

	"github.com/adityaeka26/golang-microservices/user/helper"
	"github.com/adityaeka26/golang-microservices/user/module/model/web"
	"github.com/adityaeka26/golang-microservices/user/module/service"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(c *gin.Context)
	VerifyRegister(c *gin.Context)
}
type HandlerImpl struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &HandlerImpl{
		service: service,
	}
}

func (handler *HandlerImpl) Register(c *gin.Context) {
	request := &web.RegisterRequest{}

	if err := c.ShouldBind(request); err != nil {
		helper.RespError(c, helper.CustomError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := helper.Validate(request); err != nil {
		helper.RespError(c, helper.CustomError(http.StatusBadRequest, err.Error()))
		return
	}

	err := handler.service.Register(c.Request.Context(), *request)
	if err != nil {
		helper.RespError(c, err)
		return
	}
	helper.RespSuccess(c, nil, "Register success")
}

func (handler *HandlerImpl) VerifyRegister(c *gin.Context) {
	request := &web.VerifyRegisterRequest{}

	if err := c.ShouldBind(request); err != nil {
		helper.RespError(c, helper.CustomError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := helper.Validate(request); err != nil {
		helper.RespError(c, helper.CustomError(http.StatusBadRequest, err.Error()))
		return
	}

	response, err := handler.service.VerifyRegister(c.Request.Context(), *request)
	if err != nil {
		helper.RespError(c, err)
		return
	}
	helper.RespSuccess(c, response, "Register success")
}
