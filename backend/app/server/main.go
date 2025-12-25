package main

import (
	"main/config"
	"main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitFirebase()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
