package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SerializeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Print("Serialize Middleware...")
	}
}