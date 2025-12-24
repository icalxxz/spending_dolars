package handlers

import (
	"context"
	"net/http"

	"main/config"
	"main/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const basePath = "laporan_belanja"

// CREATE
func CreateLaporan(c *gin.Context) {
	var laporan models.LaporanBelanja
	if err := c.ShouldBindJSON(&laporan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	laporan.ID = uuid.New().String()
	laporan.Total = laporan.Jumlah * laporan.Harga

	ctx := context.Background()
	db, _ := config.App.Database(ctx)

	ref := db.NewRef(basePath + "/" + laporan.ID)
	ref.Set(ctx, laporan)

	c.JSON(http.StatusCreated, laporan)
}

// READ ALL
func GetAllLaporan(c *gin.Context) {
	ctx := context.Background()
	db, _ := config.App.Database(ctx)

	var data map[string]models.LaporanBelanja
	ref := db.NewRef(basePath)

	ref.Get(ctx, &data)
	c.JSON(http.StatusOK, data)
}

// READ BY ID
func GetLaporanByID(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	db, _ := config.App.Database(ctx)

	var laporan models.LaporanBelanja
	ref := db.NewRef(basePath + "/" + id)

	if err := ref.Get(ctx, &laporan); err != nil || laporan.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, laporan)
}

// UPDATE
func UpdateLaporan(c *gin.Context) {
	id := c.Param("id")

	var laporan models.LaporanBelanja
	if err := c.ShouldBindJSON(&laporan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	laporan.ID = id
	laporan.Total = laporan.Jumlah * laporan.Harga

	ctx := context.Background()
	db, _ := config.App.Database(ctx)

	ref := db.NewRef(basePath + "/" + id)
	ref.Set(ctx, laporan)

	c.JSON(http.StatusOK, laporan)
}

// DELETE
func DeleteLaporan(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	db, _ := config.App.Database(ctx)

	ref := db.NewRef(basePath + "/" + id)
	ref.Delete(ctx)

	c.JSON(http.StatusOK, gin.H{"message": "Laporan berhasil dihapus"})
}
