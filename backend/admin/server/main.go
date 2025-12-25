package main

import (
	"main/firebase"
	"main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	firebase.InitFirebase()

	r := gin.Default()
	routes.UserRoutes(r)

	r.Run(":9000")
}
