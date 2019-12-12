package main

import (
	"fmt"

	// "github.com/alactic/ministore/sharedservice/utils/connection"
	"github.com/alactic/ministore/sharedservice/utils/router"
	"gopkg.in/couchbase/gocb.v1"
)

var bucket *gocb.Bucket

// @title Swagger Example API
// @version 1.0
// @description This is a mini store application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	fmt.Println("starting application")
	// bucket = connection.Connection()
	 router.Router()
}
