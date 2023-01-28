package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Print("Auth Middleware...")
	}
}