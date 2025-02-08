package helpers

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func SetupRedis() { 
	client := redis.NewClient(&redis.Options{ 
		Addr:     GetEnv("REDIS_HOST", "localhost:6379"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil { 
		Logger.Fatal("Redis connection failed : ", err)
		return
	}
	Logger.Info("Redis connection Ping : ", ping)

	RedisClient = client
}