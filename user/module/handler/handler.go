package handler

import (
	"net/http"

	"github.com/adityaeka26/golang-microservices/user/helper"
	"github.com/adityaeka26/golang-microservices/user/logger"
	"github.com/adityaeka26/golang-microservices/user/module/model/web"
	"github.com/adityaeka26/golang-microservices/user/module/service"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmzap/v2"
)

type Handler interface {
	Register(c *gin.Context)
	VerifyRegister(c *gin.Context)
}
type HandlerImpl struct {
	service service.Service
	logger  logger.Logger
}

func NewHandler(service service.Service, logger logger.Logger) Handler {
	return &HandlerImpl{
		service: service,
		logger:  logger,
	}
}

func (handler HandlerImpl) Register(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "Register", "handler")
	defer span.End()

	traceContextFields := apmzap.TraceContext(c.Request.Context())
	handler.logger.GetLogger().With(traceContextFields...).Debug("handling request")

	request := &web.RegisterRequest{}

	if err := c.ShouldBind(request); err != nil {
		helper.RespError(c, helper.CustomError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := helper.Validate(request); err != nil {
		helper.RespError(c, helper.CustomError(http.StatusBadRequest, err.Error()))
		return
	}

	err := handler.service.Register(ctx, *request)
	if err != nil {
		helper.RespError(c, err)
		return
	}
	helper.RespSuccess(c, nil, "Register success")
}

func (handler HandlerImpl) VerifyRegister(c *gin.Context) {
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
