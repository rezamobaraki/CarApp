package routers

import (
	"github.com/MrRezoo/CarApp/api/handlers"
	"github.com/MrRezoo/CarApp/api/middlewares"
	"github.com/MrRezoo/CarApp/config"
	"github.com/gin-gonic/gin"
)

func Users(router *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewUserHandler(cfg)
	router.POST("/send-otp/", middlewares.OTPLimiter(cfg), handler.SendOTP)
}
