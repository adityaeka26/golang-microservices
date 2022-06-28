package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(c *gin.Context)
	VerifyRegister(c *gin.Context)
}
