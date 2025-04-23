package middleware

import (
	utils "jwt-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(requestContext *gin.Context) {
		tokenString := requestContext.GetHeader("Authorization")
		if tokenString == "" {
			requestContext.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			requestContext.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			requestContext.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			requestContext.Abort()
			return
		}

		requestContext.Set("username", claims.Subject)
		requestContext.Next()
	}
}
