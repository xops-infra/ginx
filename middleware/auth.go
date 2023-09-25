package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware is a middleware function that checks for a valid token in the Authorization header
func TokenAuthMiddleware(ignorePaths []string, secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// let the request pass if request method is OPTIONS
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		// Check if the current path is in the ignore list
		for _, path := range ignorePaths {
			if strings.HasPrefix(c.Request.URL.Path, path) {
				c.Next()
				return
			}
		}

		// Check if the Authorization header is present
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing Authorization header"})
			return
		}

		// Check if the token is valid
		err := ValidateToken(token, secret)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		// Token is valid, continue
		c.Next()
	}
}

// ValidateToken
func ValidateToken(signedToken string, secret []byte) error {
	// Parse the signed JWT token and validate it
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return err
	}

	// Check if the token is expired
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			return errors.New("token expired")
		}
	}
	return nil
}
