package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"main/firebase"
	"main/models"
)

func EnsureUserExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetString("uid")
		if uid == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		ctx := context.Background()

		// 1️⃣ Cek user di Firebase Auth
		_, err := firebase.AuthClient.GetUser(ctx, uid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User not registered in Auth",
			})
			c.Abort()
			return
		}

		// 2️⃣ Cek data user di Firestore
		doc, err := firebase.FirestoreClient.
			Collection("users").
			Doc(uid).
			Get(ctx)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "User data not found in Firestore",
			})
			c.Abort()
			return
		}

		var user models.User
		if err := doc.DataTo(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to parse user data",
			})
			c.Abort()
			return
		}

		// 3️⃣ Simpan user ke context (opsional tapi sangat berguna)
		c.Set("user", user)

		c.Next()
	}
}
