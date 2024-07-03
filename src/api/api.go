package api

import (
	"github.com/MrRezoo/CarApp/src/api/routers"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	v1 := engine.Group("/api/v1/")

	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	engine.Run(":5005")
}
