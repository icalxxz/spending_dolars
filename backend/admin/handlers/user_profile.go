package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"main/firebase"
	"main/models"
)

func GetProfile(c *gin.Context) {
	uid := c.GetString("uid")

	ctx := context.Background()
	doc, err := firebase.FirestoreClient.
		Collection("users").
		Doc(uid).
		Get(ctx)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	var user models.User
	_ = doc.DataTo(&user)

	c.JSON(http.StatusOK, user)
}
