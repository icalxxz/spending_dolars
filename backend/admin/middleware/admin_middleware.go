package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"main/firebase"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		uid := c.GetString("uid")
		if uid == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		ctx := context.Background()
		doc, err := firebase.FirestoreClient.
			Collection("users").
			Doc(uid).
			Get(ctx)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		role, ok := doc.Data()["role"].(string)
		if !ok || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin access only"})
			c.Abort()
			return
		}

		c.Next()
	}
}
