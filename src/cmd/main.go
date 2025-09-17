package main

import (
	"github.com/rezamobaraki/CarApp/api"
	"github.com/rezamobaraki/CarApp/config"
	"github.com/rezamobaraki/CarApp/data/cache"
	"github.com/rezamobaraki/CarApp/data/db"
	"github.com/rezamobaraki/CarApp/data/db/migrations"
	"github.com/rezamobaraki/CarApp/pkg/logging"
)

// @BasePath /v1
// @title CarApp API
// @description This is a sample Golang server CarApp server with gin
// @contact.name Reza Mobaraki
// @contact.email rezam578@gmail.com
// @contact.url https://linkedin.com/in/reza-mobaraki
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
