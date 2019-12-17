package jwt

import (
	"fmt"

	proto "github.com/alactic/ministore/proto/userdetail"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func DecodeJWT(ctx *gin.Context, token string) error {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewUserServiceClient(conn)

	req := &proto.Requesttoken{Token: token}
	if response, err := client.UserAuthorization(ctx, req); err == nil {
		fmt.Sprint(response)
		return nil
	} else {
		return err
	}
}
