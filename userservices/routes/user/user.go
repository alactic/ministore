package user

import (
	"github.com/alactic/ministore/userservices/controllers/user"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	users := router.Group("/users") 
		{
			users.POST("/", user.CreateUserEndpoint)
		}
}
