package user

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/alactic/ministore/userservices/models/userm"
	"github.com/alactic/ministore/userservices/utils/connection"

	hashed "github.com/alactic/ministore/userservice/utils/hash"
	"github.com/alactic/ministore/userservices/utils/shared/error"

	// "github.com/alactic/ministore/sharedservice/httputil"

	"gopkg.in/couchbase/gocb.v1"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var bucket *gocb.Bucket = connection.Connection()

// CreateUserEndpoint godoc
// @Summary Add a user
// @Description add by json user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body userm.User true "User Data"
// @Success 200 {object} userm.User
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/users [post]
func CreateUserEndpoint(ctx *gin.Context) {
	var addUser userm.User
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) == 0 {
		error.NewError(ctx, http.StatusBadRequest, errors.New("please set Header Authorization"))
		return
	}
	e := ctx.ShouldBindJSON(&addUser)
	if e != nil || e == nil {
		if err := userm.Validation(addUser); err != nil {
			error.NewError(ctx, http.StatusBadRequest, err)
			return
		}
	}

	addUser.ID = uuid.Must(uuid.NewV4()).String()
	addUser.Password = string(hashed.Hash(addUser.Password))
	addUser.Confirmpassword = ""
	_, err := bucket.Insert(addUser.ID, addUser, 0)
	if err != nil {
		fmt.Println("error :: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	response := make(map[string]string)
	response["Firstname"] = addUser.Firstname
	response["Lastname"] = addUser.Lastname
	response["Type"] = addUser.Type
	response["Email"] = addUser.Email
	ctx.JSON(http.StatusOK, response)
}

// UpdateUserEndpoint godoc
// @Summary Update a user
// @Description Update by json user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body userm.UpdateUser true "User Data"
// @Success 200 {object} userm.UpdateUser
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/users [put]
func UpdateUserEndpoint(ctx *gin.Context) {
	var updateUser userm.UpdateUser
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) == 0 {
		error.NewError(ctx, http.StatusBadRequest, errors.New("please set Header Authorization"))
		return
	}
	if err := userm.ValidationUpdate(updateUser); err != nil {
		error.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err := bucket.Replace(updateUser.ID, updateUser, 0, 0)
	if err != nil {
		fmt.Println("error :: ", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, updateUser)
}

// GetUsersEndpoint godoc
// @Summary get all users
// @Description get all user by json
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} userm.UpdateUser
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/users [get]
func GetUsersEndpoint(ctx *gin.Context) {
	var users []userm.UpdateUser
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) == 0 {
		error.NewError(ctx, http.StatusBadRequest, errors.New("please set Header Authorization"))
		return
	}

	query := gocb.NewN1qlQuery("SELECT " + bucket.Name() + ".* FROM " + bucket.Name())
	rows, err := bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var row userm.UpdateUser
	for rows.Next(&row) {
		users = append(users, row)
	}

	var response = make(map[string][]userm.UpdateUser)
	response["data"] = users
	ctx.JSON(http.StatusOK, response)
}

// GetUserById godoc
// @Summary Show a user
// @Description get user by ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} userm.UpdateUser
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Router /api/v1/users/{id} [get]
func GetUserById(ctx *gin.Context) {
	var user userm.UpdateUser
	id := ctx.Param("id")
	_, err := bucket.Get(id, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, user)
}

// TestEndpoint godoc
// @Summary testing endpoint
// @Description testing endpoint
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} userm.UpdateUser
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/users/test [get]
func TestEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "testing app " + os.Getenv("MODULE_NAME")})
}
