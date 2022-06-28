package middleware

import "github.com/gin-gonic/gin"

type Auth interface {
	VerifyJWTToken() gin.HandlerFunc
}
