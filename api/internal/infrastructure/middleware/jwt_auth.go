package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	authInfra "github.com/hryt430/RESTAPI/api/internal/infrastructure/auth"
)

func AuthMiddleware(jwtRepo *authInfra.JwtAuthRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		user, err := jwtRepo.Validate(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// コンテキストにユーザー情報を保持
		ctx.Set("user", user)
		ctx.Next()
	}
}
