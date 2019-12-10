package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alactic/ministore/sharedservice/utils/shared/error"
	_ "github.com/alactic/ministore/userservices/docs"
	"github.com/alactic/ministore/userservices/routes/index"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"context"
	"github.com/alactic/ministore/userservices/proto/userdetails"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Router() {
	r := gin.Default()
  
	v1 := r.Group("/api/v1")
	{
		index.Index(v1)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8888")

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(r); e != nil {
		panic(e)
	}
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			error.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}


func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}
