package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"main/firebase"
)

func ValidateUser(c *gin.Context) {
	uid := c.GetString("uid")
	if uid == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"valid": false})
		return
	}

	ctx := context.Background()
	_, err := firebase.FirestoreClient.
		Collection("users").
		Doc(uid).
		Get(ctx)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"registered": false,
			"reason":     "not_registered",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"registered": true,
	})
}
