package router

import (
	"github.com/gin-gonic/gin"
	// "github.com/swaggo/gin-swagger/swaggerFiles"

	// "github.com/swaggo/gin-swagger/swaggerFiles"

	// "github.com/alactic/ministore/shareservice/utils/shared/error"
	// "github.com/alactic/ministore/authservices/routes/index"
	_ "github.com/alactic/ministore/userservices/docs"
	"github.com/alactic/ministore/userservices/routes/index"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct{}

func Router() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		index.Index(v1)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":1111")
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
