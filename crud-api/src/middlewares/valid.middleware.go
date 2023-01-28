package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ValidMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Print("Valid Middleware...")	
	}
}