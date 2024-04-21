package helpers

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/RLungWu/MyShortURL/db"

	"github.com/go-redis/redis"
)

func CheckRateLimit(ip string) error {
	r := db.CreateClient(0)
	defer r.Close()

	val, err := r.Get(db.Context, ip).Result()
	if err == redis.Nil {
		quota, err := strconv.Atoi("10")
		if err != nil {
			return errors.New("invalid API Quota")
		}
		_ = r.Set(db.Context, ip, quota, 30*time.Minute).Err()
	} else if err != nil {
		return errors.New("cannot connect to server")
	} else {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			return errors.New("invalid rate limit")
		}
		if valInt <= 0 {
			limit, _ := r.TTL(db.Context, ip).Result()
			return fmt.Errorf("rate limit exceeded, try again in %v", limit/time.Second)
		}

		if err := r.Decr(db.Context, ip).Err(); err != nil {
			return errors.New("failed to decrement quota: " + err.Error())
		}
	}

	return nil
}
