package auth

import (
	"fmt"
	"net/http"

	"github.com/alactic/ministore/authservices/models/userm"
	"github.com/alactic/ministore/authservices/utils/connection"

	hashed "github.com/alactic/ministore/authservices/utils/hash"
	jwtFile "github.com/alactic/ministore/authservices/utils/jwt"

	"gopkg.in/couchbase/gocb.v1"

	proto "github.com/alactic/ministore/proto/userdetail"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var bucket *gocb.Bucket = connection.Connection()

// CreateAuthEndpoint godoc
// @Summary authenticate user
// @Description authenticate user by providing username and password
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param user body userm.User true "User Data"
// @Success 200 {object} userm.User
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/auth/login [post]
func CreateAuthEndpoint(ctx *gin.Context) {
	var user userm.User
	var users []userm.UserDetails
	e := ctx.ShouldBindJSON(&user)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, e.Error())
		return
	}
	query := gocb.NewN1qlQuery("SELECT " + bucket.Name() + ".* FROM " + bucket.Name() + " WHERE `email` = $1")
	rows, err := bucket.ExecuteN1qlQuery(query, []interface{}{user.Email})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var row userm.UserDetails
	for rows.Next(&row) {
		users = append(users, row)
	}
	error := hashed.CompareHashValue(user.Password, users[0].Password)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid username or password"})
		return
	}
	details := make(map[string]string)
	details["firstname"] = users[0].Firstname
	details["lastname"] = users[0].Lastname
	details["email"] = users[0].Email

	token, err := jwtFile.GenerateJWT(details)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": err})
	} else {
		users[0].Token = token
		var response = make(map[string][]userm.UserDetails)
		response["data"] = users

		ctx.JSON(http.StatusOK, gin.H{"data": response})
	}
}

// VerifyByEmail godoc
// @Summary verify grpc
// @Description Verify verify grpc
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param user body userm.UserEmail true "User Data"
// @Success 200 {object} userm.User
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/auth/verify [post]
func VerifyByEmail(ctx *gin.Context) {
	var user userm.UserEmail
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewUserServiceClient(conn)

	req := &proto.Request{Email: user.Email}
	if response, err := client.UserDetails(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result details": fmt.Sprint(response),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// VerifyByEmail godoc
// @Summary testing grpc
// @Description testing grpc
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Success 200 {object} userm.User
// @Failure 400 {object} error.HTTPError
// @Failure 404 {object} error.HTTPError
// @Failure 500 {object} error.HTTPError
// @Security ApiKeyAuth
// @Router /api/v1/auth/test [get]
func TestByEmail(ctx *gin.Context) {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewUserServiceClient(conn)

	req := &proto.Request{Email: "elvis@lendsqr.com"}
	if response, err := client.UserDetails(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result details docker 2334": fmt.Sprint(response),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
