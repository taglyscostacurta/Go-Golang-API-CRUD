package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taglyscostacurta/Go-Golang-API-CRUD/routes"
)

func main() {
	service := gin.Default()

	routes.GetRoutes(service)

	service.Run()
}
