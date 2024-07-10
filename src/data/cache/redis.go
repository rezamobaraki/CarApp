package cache

import (
	"context"
	"github.com/MrRezoo/CarApp/config"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var redisClient *redis.Client

func InitRedis(config *config.Config) {
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis at %s:%s, error: %v", config.Redis.Host, config.Redis.Port, err)
	} else {
		log.Printf("Redis connected to %s:%s", config.Redis.Host, config.Redis.Port)
	}

}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	_ = redisClient.Close()
}
