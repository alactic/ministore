package router

import (
	"context"
	"net"

	"github.com/gin-gonic/gin"
	_ "github.com/alactic/ministore/shareservice/docs"

	proto "github.com/alactic/ministore/proto/userdetail"

	"github.com/alactic/ministore/shareservice/controllers/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func Router() {
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
	var userdetail = user.UserDetails(request.GetEmail())

	return &proto.Response{
		Firstname: userdetail["email"],
		Lastname:  userdetail["lastname"],
		Email:     userdetail["firstname"]}, nil
}
