package router

import (
	"context"
	"fmt"
	"net"

	"github.com/gin-gonic/gin"

	// "github.com/alactic/ministore/sharedservice/utils/shared/error"
	_ "github.com/alactic/ministore/userservices/docs"

	proto "github.com/alactic/ministore/proto/userdetail"

	"github.com/alactic/ministore/userservices/controllers/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func Router() {
	// r := gin.Default()

	// v1 := r.Group("/api/v1")
	// {
	// 	index.Index(v1)
	// }
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.Run(":8888")

	lis, _ := net.Listen("tcp", ":50051")
	srv := grpc.NewServer()
	proto.RegisterUserServiceServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(lis); e != nil {
		panic(e)
	}
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			// error.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}

func (s *server) UserDetails(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	fmt.Println("user details emial :: ", request.GetEmail())
	var userdetail = user.UserDetails(request.GetEmail())

	return &proto.Response{
		Firstname: userdetail.firstname,
		Lastname:  userdetail.lastname,
		Email:     userdetail.email}, nil
}
