package helpers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"
	"trunc-it/trunc.it/redirector/config"

	"github.com/redis/go-redis/v9"
)

type cache struct{}

var (
	Cache *cache
)

func (c *cache) Store(key any, document any, ttl time.Duration) (string, error) {
	r := config.GetRedis()

	stringifiedKey := fmt.Sprintf("%v", key)

	h := sha256.New()
	h.Write([]byte(stringifiedKey))

	hashKey := h.Sum(nil)

	encodedHashKey := hex.EncodeToString(hashKey)
	err := r.Set(context.Background(), encodedHashKey, document, ttl).Err()

	if err != nil {
		return "", fmt.Errorf("Failed to store document: %v", err)
	}

	return encodedHashKey, nil
}

func (c *cache) Lookup(key any) (*string, error) {
	r := config.GetRedis()

	stringifiedKey := fmt.Sprintf("%v", key)

	h := sha256.New()
	h.Write([]byte(stringifiedKey))

	hashKey := h.Sum(nil)

	data, err := r.Get(context.Background(), string(hashKey)).Result()

	log.Printf("Error during cache lookup: %v", err)

	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to get data: %v", err)
	}

	return &data, nil
}
