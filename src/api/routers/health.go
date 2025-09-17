package routers

import (
	"github.com/rezamobaraki/CarApp/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(router *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	router.GET("/", handler.Health)
}
