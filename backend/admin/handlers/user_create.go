package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"main/firebase"
	"main/models"
)

type RegisterRequest struct {
	Name string `json:"name" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	uid := c.GetString("uid")
	email := c.GetString("email")

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	ctx := context.Background()
	docRef := firebase.FirestoreClient.Collection("users").Doc(uid)

	// cek apakah sudah terdaftar
	_, err := docRef.Get(ctx)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already registered"})
		return
	}

	user := models.User{
		UID:       uid,
		Email:     email,
		Name:      req.Name,
		Role:      "user",
		CreatedAt: time.Now(),
	}

	if _, err := docRef.Set(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "registration success",
	})
}
