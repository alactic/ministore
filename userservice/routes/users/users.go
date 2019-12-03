package users

import (
	"github.com/alactic/ministore/userservice/controllers/users"
	"github.com/alactic/ministore/userservice/middlewares/authentication"
	"github.com/gorilla/mux"
)

func Users(router *gin.Context) {
	users := router.Group("/users")
		{
			users.POST("/", users.CreateUserEndpoint)
		}
}
