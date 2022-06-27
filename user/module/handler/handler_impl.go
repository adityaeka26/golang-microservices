package handler

import (
	"net/http"

	"github.com/adityaeka26/golang-microservices/user/helper"
	"github.com/adityaeka26/golang-microservices/user/module/model/web"
	"github.com/adityaeka26/golang-microservices/user/module/service"
	"github.com/gin-gonic/gin"
)

type HandlerImpl struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &HandlerImpl{
		service: service,
	}
}

func (handler *HandlerImpl) GetUser(c *gin.Context) {
	request := &web.GetUserRequest{}

	if err := c.ShouldBindQuery(request); err != nil {
		helper.RespError(c, helper.CustomError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := helper.Validate(request); err != nil {
		helper.RespError(c, helper.CustomError(http.StatusBadRequest, err.Error()))
		return
	}

	response, err := handler.service.GetUser(c.Request.Context(), *request)
	if err != nil {
		helper.RespError(c, err)
		return
	}
	helper.RespSuccess(c, response, "Get user success")
}

func (handler *HandlerImpl) CreateUser(c *gin.Context) {
	request := &web.CreateUserRequest{}

	if err := c.ShouldBind(request); err != nil {
		helper.RespError(c, helper.CustomError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := helper.Validate(request); err != nil {
		helper.RespError(c, helper.CustomError(http.StatusBadRequest, err.Error()))
		return
	}

	response, err := handler.service.CreateUser(c.Request.Context(), *request)
	if err != nil {
		helper.RespError(c, err)
		return
	}
	helper.RespSuccess(c, response, "Create user success")
}
