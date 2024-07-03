package routers

import (
	"github.com/MrRezoo/CarApp/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(router *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	router.GET("/", handler.Health)
}
