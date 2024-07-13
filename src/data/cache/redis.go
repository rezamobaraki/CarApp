package cache

import (
	"github.com/MrRezoo/CarApp/config"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var redisClient *redis.Client

func InitRedis(config *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:               config.Redis.Host + ":" + config.Redis.Port,
		Password:           config.Redis.Password,
		DB:                 config.Redis.DB,
		DialTimeout:        config.Redis.DialTimeout * time.Second,
		ReadTimeout:        config.Redis.ReadTimeout * time.Second,
		WriteTimeout:       config.Redis.WriteTimeout * time.Second,
		PoolSize:           config.Redis.PoolSize,
		PoolTimeout:        config.Redis.PoolTimeout * time.Second,
		IdleTimeout:        config.Redis.IdleTimeout * time.Millisecond,
		IdleCheckFrequency: config.Redis.IdleCheckFrequency * time.Millisecond,
	})

	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		return err
	}

	log.Printf("Redis connected to %s:%s", config.Redis.Host, config.Redis.Port)
	return nil

}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	_ = redisClient.Close()
}
