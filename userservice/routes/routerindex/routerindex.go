package routerindex

import (
	"github.com/alactic/ministore/userservice/routes/users"
	"github.com/gin-gonic/gin"
)

func Routerindex(router *gin.RouterGroup) {
	users.Users(router)
}
