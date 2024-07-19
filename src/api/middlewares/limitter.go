package middlewares

import (
	"github.com/MrRezoo/CarApp/api/helper"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LimitByRequestCount() gin.HandlerFunc {
	limiter := tollbooth.NewLimiter(1, nil)
	return func(context *gin.Context) {
		httpError := tollbooth.LimitByRequest(limiter, context.Writer, context.Request)
		if httpError != nil {
			context.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(nil, false, -100, httpError),
			)
			return
		} else {
			context.Next()
		}

	}
}
