package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	hh "github.com/patsnapops/http-headers"
	"github.com/spf13/cast"
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

		token, err := extractBearerToken(c.Request.Header)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		claims, err := parse(token, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		// log.Debugf(tea.Prettify(claims))
		err = checkIfExpired(claims)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		c.Next()
	}
}

func extractBearerToken(header http.Header) (string, error) {
	const BearerSpace = "Bearer "
	auth := header.Get(hh.Authorization)
	token := strings.TrimPrefix(auth, BearerSpace)
	if token == auth {
		return "", errors.New("no bearer token found in Authorization header")
	}
	return token, nil
}

func checkIfExpired(claims map[string]any) error {
	const Exp = "exp"
	exp := claims[Exp]
	if exp == nil {
		return errors.New("exp not found")
	}
	expTimestamp, ok := exp.(float64)
	if !ok {
		return errors.New("exp is not a valid int64")
	}
	if time.Now().Unix() > cast.ToInt64(expTimestamp) {
		return errors.New("token expired")
	}
	return nil
}

// parse use HMAC to decrypt the token
func parse(token string, secret []byte) (map[string]any, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			const Alg = "alg"
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header[Alg])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
