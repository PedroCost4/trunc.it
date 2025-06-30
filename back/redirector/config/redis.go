package config

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
)

func SetupRedis() error {
	addr := os.Getenv("REDIS_ADDRESS")
	password := os.Getenv("REDIS_PASSWORD")

	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	err := client.Ping(context.Background()).Err()

	if err != nil {
		return fmt.Errorf("Failed to ping redis: %v", err)
	}

	return nil
}

func GetRedis() *redis.Client {
	return client
}
