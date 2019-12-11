package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	// "github.com/alactic/ministore/sharedservice/utils/shared/error"
	_ "github.com/alactic/ministore/userservices/docs"
	"github.com/alactic/ministore/userservices/routes/index"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	r := gin.Default()
  
	v1 := r.Group("/api/v1")
	{
		index.Index(v1)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8888")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authorization testing")
		if len(c.GetHeader("Authorization")) == 0 {
			// error.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
