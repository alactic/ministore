package index

import (
	"github.com/alactic/ministore/shareservice/routes/user"
	"github.com/gin-gonic/gin"
)

func Index(router *gin.RouterGroup) {
	user.User(router)
}
