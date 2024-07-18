package cache

import (
	"encoding/json"
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/pkg/logging"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisClient *redis.Client
var logger = logging.NewLogger(config.GetConfig())

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

	logger.Info(logging.Redis, logging.Startup, "Redis connected", nil)
	return nil

}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	_ = redisClient.Close()
}

func Set[T any](c *redis.Client, key string, value T, expiration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(c.Context(), key, v, expiration).Err()
}

func Get[T any](c *redis.Client, key string) (T, error) {
	var dest = *new(T)
	v, err := c.Get(c.Context(), key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(v), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil

}
