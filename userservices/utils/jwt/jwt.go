package jwt

import (
	proto "github.com/alactic/ministore/proto/userdetail"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func DecodeJWT(ctx *gin.Context, token string) int64 {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewUserServiceClient(conn)

	req := &proto.Requesttoken{Token: token}
	if response, error := client.UserAuthorization(ctx, req); error == nil {
		return int64(response.Authorization)
	} 	 
	return 400
}
