package main

import (
	"github.com/MrRezoo/CarApp/api"
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/data/cache"
	"github.com/MrRezoo/CarApp/data/db"
	"log"
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

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatalln("Failed to connect to Redis", err)
	}

	err = db.InitDB(cfg)
	defer db.CloseDB()
	if err != nil {
		log.Fatalln("Failed to connect to Postgres", err)
	}
	api.InitServer(cfg)
}
