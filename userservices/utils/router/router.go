package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alactic/ministore/sharedservice/httputil"
	// _ "github.com/alactic/ministore/userservices/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	r := gin.Default()

	// v1 := r.Group("/api/v1")
	// {
	// 	index.Index(v1)
	// }
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8088")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
