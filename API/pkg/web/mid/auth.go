package mid

import (
	"fmt"
	"hospitalApi/cmd/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format. Must be Bearer <token>"})
			return
		}

		token := parts[1]

		isValid, username, err := validateToken(token)
		if err != nil {
			fmt.Printf("Token validation error: %v\n", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error during token validation"})
			return
		}

		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Request.Header.Set("username", username)

		c.Next()
	}
}

func validateToken(tokenString string) (isValid bool, username string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		configuration := config.New()
		jwtSecretBytes := []byte(configuration.SecretKey)
		return jwtSecretBytes, nil
	})

	if err != nil {
		return false, "", fmt.Errorf("Invalid token: %v", err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return false, "", fmt.Errorf("Failed to extract username from token")
		}

		return true, username, nil

	} else {
		return false, "", fmt.Errorf("Invalid token")
	}

}
