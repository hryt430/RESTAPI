package jwt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_auth "github.com/hryt430/RESTAPI/api/internal/infrastructure/auth"
)

func AuthMiddleware(jwtRepo *jwt_auth.JwtAuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		user, err := jwtRepo.Validate(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// コンテキストにユーザー情報を保持
		c.Set("user", user)
		c.Next()
	}
}
