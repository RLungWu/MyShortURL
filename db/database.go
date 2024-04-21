package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Context = context.Background()

func CreateClient(dbno int) *redis.Client{
	// CreateClient creates a new redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbno,
	})

	return rdb
}
