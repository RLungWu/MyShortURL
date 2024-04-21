package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Context = context.Background()

func CreateClient(dbno int) *redis.Client {
	// CreateClient creates a new redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       dbno,
	})

	return rdb
}
