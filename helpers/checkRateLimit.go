package helpers

import (
    "errors"
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
        quota := 1000
        _ = r.Set(db.Context, ip, quota-1, 30*time.Minute).Err()
        return nil 
    } else if err != nil {
        return err
    } else {
        valInt, err := strconv.Atoi(val)
        if err != nil {
            return err 
        }
        if valInt <= 0 {
            return errors.New("rate limit exceeded")
        }
        _ = r.Decr(db.Context, ip).Err() 
    }

    return nil
}
