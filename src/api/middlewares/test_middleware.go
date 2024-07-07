package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		apiKey := context.GetHeader("x-api-key")
		if apiKey == "1" {
			context.Next()
		}
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized Access!",
		})
		return
	}
}
