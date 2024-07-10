package api

import (
	"fmt"
	"github.com/MrRezoo/CarApp/api/middlewares"
	"github.com/MrRezoo/CarApp/api/routers"
	"github.com/MrRezoo/CarApp/api/validations"
	"github.com/MrRezoo/CarApp/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()
	engine := gin.New()
	engine.Use(middlewares.Cors(cfg))
	engine.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.LimitByRequestCount())
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.ValidateIranianMobile, true)
		val.RegisterValidation("password", validations.ValidatePassword, true)
	}

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
