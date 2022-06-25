package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	GetUser(c *gin.Context)
}
