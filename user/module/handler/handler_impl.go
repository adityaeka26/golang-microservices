package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerImpl struct{}

func NewHandler() Handler {
	return &HandlerImpl{}
}

func (handler *HandlerImpl) GetUser(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
