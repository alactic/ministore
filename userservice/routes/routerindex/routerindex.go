package routerindex

import (
	"github.com/alactic/ministore/userservice/routes/users"
	"github.com/gorilla/mux"
)

func Routerindex(router *gin.Context) {
	users.users(router)
}
