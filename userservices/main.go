package main

import (
	"fmt"

	"github.com/alactic/ministore/userservice/utils/router"
)

func main() {
	fmt.Println("starting application")
	router.Router()
}
