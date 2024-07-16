package main

import (
	"github.com/MrRezoo/CarApp/api"
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/data/cache"
	"github.com/MrRezoo/CarApp/data/db"
	"github.com/MrRezoo/CarApp/data/db/migrations"
	"github.com/MrRezoo/CarApp/pkg/logging"
)

// @BasePath /v1
// @title CarApp API
// @description This is a sample Golang server CarApp server with gin
// @contact.name Reza Mobaraki
// @contact.email rezam578@gmail.com
// @contact.url http://www.rezoo.ir
// @version 1.0
// @schemes http
// @host localhost:5005
// @BasePath /api/
// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDB(cfg)
	defer db.CloseDB()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migrations.Up()

	api.InitServer(cfg)
}
