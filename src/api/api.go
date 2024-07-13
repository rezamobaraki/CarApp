package api

import (
	"fmt"
	"github.com/MrRezoo/CarApp/api/middlewares"
	"github.com/MrRezoo/CarApp/api/routers"
	"github.com/MrRezoo/CarApp/api/validations"
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	engine := gin.New()

	RegisterValidators()
	RegisterMiddlewares(engine, cfg)
	RegisterRoutes(engine)
	RegisterSwagger(engine, cfg)

	engine.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(engine *gin.Engine) {
	api := engine.Group("/api/")
	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health/")
		test := v1.Group("/test/")
		routers.Health(health)
		routers.TestRouter(test)
	}

	v2 := api.Group("/v2")
	{
		healthRouter := v2.Group("/health")
		routers.Health(healthRouter)
	}

}

func RegisterMiddlewares(engine *gin.Engine, cfg *config.Config) {
	engine.Use(middlewares.Cors(cfg))
	engine.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.LimitByRequestCount())
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.ValidateIranianMobile, true)
		val.RegisterValidation("password", validations.ValidatePassword, true)
	}
}

func RegisterSwagger(engine *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "CarApp API"
	docs.SwaggerInfo.Description = "This is a sample Golang server CarApp server with gin"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = "localhost" + ":" + cfg.Server.Port
	docs.SwaggerInfo.Schemes = []string{"http"}
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
