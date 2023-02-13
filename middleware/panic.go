package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// TODO: support custom log facility
				log.Printf("panic: %v\n", r)
			}
		}()

		c.Next()
	}
}
