package users

import (
	"net/http"
	"time"

	"github.com/alactic/ministore/userservice/models/users"
	"github.com/alactic/ministore/userservice/utils/connection"
	"github.com/alactic/ministore/sharedservice/httputil"

	"gopkg.in/couchbase/gocb.v1"

	"github.com/gin-gonic/gin"
)

var bucket *gocb.Bucket = connection.Connection()

// CreateUserEndpoint godoc
// @Summary Add a account
// @Description add by json account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body users.Customer true "Add account"
// @Success 200 {object} users.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [post]
func CreateUserEndpoint(ctx *gin.Context) {
	var addAccount users.Customer
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := addAccount.Validation(); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := users.Customer{
		Name: addAccount.Name,
	}
	lastID, err := account.Insert()
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account.ID = lastID
	ctx.JSON(http.StatusOK, account)
}