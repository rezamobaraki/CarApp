package api

import (
	"fmt"
	"github.com/MrRezoo/CarApp/api/routers"
	"github.com/MrRezoo/CarApp/config"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	v1 := engine.Group("/api/v1/")

	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	engine.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
