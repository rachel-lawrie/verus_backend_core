package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rachel-lawrie/verus_backend_core/common"
	"github.com/rachel-lawrie/verus_backend_core/utils"
)

func CombinedAuthMiddleware(collection common.CollectionInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims := &utils.Claims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return utils.JwtKey, nil
			})

			if err == nil && token.Valid {
				c.Set("cockpit_user_id", claims.UserID)
				c.Next()
				return
			}
		}

		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		hashedKey := utils.HashAPIKey(apiKey)
		var secret struct {
			ClientID string `bson:"client_id"`
		}
		err := collection.FindOne(context.Background(), map[string]interface{}{
			"client_secret_hash": hashedKey,
			"revoked":            false,
			"deleted_at":         nil,
		}).Decode(&secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or inactive API key"})
			c.Abort()
			return
		}

		c.Set("client_id", secret.ClientID)
		c.Next()
	}
}
