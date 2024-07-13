package main

import (
	"github.com/MrRezoo/CarApp/api"
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/data/cache"
	"github.com/MrRezoo/CarApp/data/db"
	"log"
)

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
