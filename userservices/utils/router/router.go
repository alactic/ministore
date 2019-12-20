package router

import (
	"net/http"
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
	r.Run(":2222")
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		error := jwtFile.DecodeJWT(c, token)
		if error != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorised to perform this action"})
			c.Abort()
			return
		}
		c.Next()
	}
}

