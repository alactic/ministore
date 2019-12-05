package index

import (
	"github.com/alactic/ministore/authservices/routes/auth"
	"github.com/gin-gonic/gin"
)

func Index(router *gin.RouterGroup) {
	auth.Auth(router)
}
