package router

import (
	"net"

	"github.com/gin-gonic/gin"

	// "github.com/alactic/ministore/sharedservice/utils/shared/error"
	_ "github.com/alactic/ministore/userservices/docs"

	"context"

	proto "github.com/alactic/ministore/userservices/proto/userdetails"

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
	request.GetEmail()

	result := make(map[string]string)
	result["firstname"] = "elvis"
	result["lastname"] = "okafor"
	result["email"] = "elvis@lendsqr.com"

	return &proto.Response{Firstname: "elvis", Lastname: "okafor", Email: "elvis@gmail.com"}, nil
}