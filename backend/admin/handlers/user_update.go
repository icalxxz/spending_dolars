package handlers

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"

	"main/firebase"
)

func UpdateUser(c *gin.Context) {
	uid := c.Param("uid")

	var input struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Role  string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	// Update Firebase Auth
	_, err := firebase.AuthClient.UpdateUser(ctx, uid, (&auth.UserToUpdate{}).
		Email(input.Email),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update Firestore
	_, err = firebase.FirestoreClient.
		Collection("users").
		Doc(uid).
		Update(ctx, []firestore.Update{
			{Path: "email", Value: input.Email},
			{Path: "name", Value: input.Name},
			{Path: "role", Value: input.Role},
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}
