package routes

import (
	"github.com/gin-gonic/gin"

	"main/handlers"
	"main/middleware"
)

func UserRoutes(r *gin.Engine) {
	api := r.Group("/api/users")
	api.Use(middleware.FirebaseAuth())

	api.POST("/register", handlers.RegisterUser)
	api.GET("/validate", handlers.ValidateUser)
	api.GET("/profile", handlers.GetProfile)
	api.PUT("/profile", handlers.UpdateProfile)

	admin := api.Group("")
	admin.Use(middleware.AdminOnly())
	{
		admin.GET("", handlers.GetUsers)
		admin.GET("/:uid", handlers.GetUserByID)
		admin.PUT("/:uid", handlers.UpdateUser)
		admin.DELETE("/:uid", handlers.DeleteUser)
	}
}
