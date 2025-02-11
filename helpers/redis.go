package helpers

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func SetupRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: GetEnv("REDIS_HOST", "localhost:6379"),
		DB:   0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		Logger.Error("failed to connect redis, ", err)
	}
	Logger.Info("PING REDIS: " + ping)

	RedisClient = client
}
