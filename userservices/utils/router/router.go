package router

import (
	_ "github.com/alactic/ministore/userservices/docs"
	"github.com/alactic/ministore/userservices/routes/index"
	jwtFile "github.com/alactic/ministore/userservices/utils/jwt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct{}

func Router() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.Use(Authorization())
	{
		index.Index(v1)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":1111")
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtFile.DecodeJWT(c, c.GetHeader("Authorization"))
	}
}
