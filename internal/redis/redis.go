package redis

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Rdb *redis.Client

func InitReddis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
}

func SetValue(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(Ctx, 5*time.Second)
	defer cancel()
	return Rdb.Set(ctx, key, value, 0).Err()
}

func GetValue(key string) (string, error) {
	ctx, cancel := context.WithTimeout(Ctx, 5*time.Second)
	defer cancel()
	val, err := Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}
