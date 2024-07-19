package middlewares

import (
	"errors"
	"github.com/MrRezoo/CarApp/api/helper"
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func OTPLimiter(cfg *config.Config) gin.HandlerFunc {
	var ipRateLimiter = limiter.NewIPRateLimiter(rate.Every(cfg.OTP.Limiter*time.Second), 1)
	return func(c *gin.Context) {
		limiter := ipRateLimiter.GetLimiter(c.Request.RemoteAddr)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, helper.OtpLimiterError, errors.New("not allowed")))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
