package routers

import (
	"github.com/rezamobaraki/CarApp/api/handlers"
	"github.com/rezamobaraki/CarApp/api/middlewares"
	"github.com/rezamobaraki/CarApp/config"
	"github.com/gin-gonic/gin"
)

func Users(router *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewUserHandler(cfg)
	router.POST("/send-otp/", middlewares.OTPLimiter(cfg), handler.SendOTP)
}
