package main

import (
	"github.com/MrRezoo/CarApp/api"
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}
