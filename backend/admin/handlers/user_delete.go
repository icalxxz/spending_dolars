package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"main/firebase"
)

func DeleteUser(c *gin.Context) {
	uid := c.Param("uid")
	ctx := context.Background()

	// Delete Firestore data
	_, err := firebase.FirestoreClient.
		Collection("users").
		Doc(uid).
		Delete(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user data",
		})
		return
	}

	// Delete Firebase Auth user
	err = firebase.AuthClient.DeleteUser(ctx, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete auth user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
