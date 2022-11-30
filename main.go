package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	service := gin.Default()

	getRoutes(service)

	service.Run()
}
