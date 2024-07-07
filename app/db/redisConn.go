package db

import (
	"context"
	"gosvr/logging"

	"github.com/go-redis/redis/v8"
)

func RedisInit() {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.20.0.4:6379", // Redis server address
		Password: "",                // No password set
		DB:       0,                 // Use the default database
	})

	// Ping the Redis server to check connectivity
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logging.Log.Error("Failed to connect to Redis: ", err)
		return
	}
	logging.Log.Info("Connected to Redis:", pong)
}
