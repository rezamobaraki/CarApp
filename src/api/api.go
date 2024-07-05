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

	api := engine.Group("/api/")
	v1 := api.Group("/v1/")

	{
		healthRouter := v1.Group("/health")
		testRouter := v1.Group("/test")
		routers.Health(healthRouter)
		routers.TestRouter(testRouter)
	}

	v2 := api.Group("/v2/")
	{
		healthRouter := v2.Group("/health")
		routers.Health(healthRouter)
	}

	err := engine.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		panic(err)
	}
}
