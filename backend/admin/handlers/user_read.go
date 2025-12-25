package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"main/firebase"
	"main/models"
)

func GetUsers(c *gin.Context) {
	ctx := context.Background()
	iter := firebase.FirestoreClient.Collection("users").Documents(ctx)

	var users []models.User
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var user models.User
		_ = doc.DataTo(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	uid := c.Param("uid")

	doc, err := firebase.FirestoreClient.
		Collection("users").
		Doc(uid).
		Get(context.Background())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var user models.User
	_ = doc.DataTo(&user)

	c.JSON(http.StatusOK, user)
}
