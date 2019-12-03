package user

import (
	"net/http"

	users "github.com/alactic/ministore/userservices/models/user"
	"github.com/alactic/ministore/userservices/utils/connection"
	// "github.com/alactic/ministore/sharedservice/httputil"

	"gopkg.in/couchbase/gocb.v1"

	"github.com/gin-gonic/gin"
)

var bucket *gocb.Bucket = connection.Connection()

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}


// CreateUserEndpoint godoc
// @Summary Add a account
// @Description add by json account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body string true "Add account"
// @Success 200 {object} string
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /accounts [post]
func CreateUserEndpoint(ctx *gin.Context) {
	var addUser users.User
	if err := ctx.ShouldBindJSON(&addUser); err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	// if err := addUser.Validation(); err != nil {
	// 	httputil.NewError(ctx, http.StatusBadRequest, err)
	// 	return
	// }
	user := users.User{
		Firstname: addUser.Firstname,
	}
	lastID, err := user.Insert()
	if err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	user.ID = lastID
	ctx.JSON(http.StatusOK, user)
}