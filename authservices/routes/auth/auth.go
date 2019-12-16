package auth

import (
	"github.com/alactic/ministore/authservices/controllers/auth"
	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup) {
	route := router.Group("/auth")
	{
		route.POST("/login", auth.CreateAuthEndpoint)
		route.POST("/verify", auth.VerifyByEmail)
		route.GET("/test", auth.TestByEmail)
	}
}
