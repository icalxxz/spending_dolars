package routes

import (
	"main/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/laporan", handlers.CreateLaporan)
		api.GET("/laporan", handlers.GetAllLaporan)
		api.GET("/laporan/:id", handlers.GetLaporanByID)
		api.PUT("/laporan/:id", handlers.UpdateLaporan)
		api.DELETE("/laporan/:id", handlers.DeleteLaporan)
	}
}
