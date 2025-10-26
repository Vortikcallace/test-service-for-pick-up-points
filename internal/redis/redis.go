package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Ctx = context.Background()
var Rdb *redis.Client

type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		Host:     v.GetString("redis.host"),
		Port:     v.GetInt("redis.port"),
		Password: v.GetString("redis.password"),
		DB:       v.GetInt("redis.db"),
	}
}

func InitRedisConfig(cfg *Config) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: 10,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis")
}

func InitRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	Rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis")
}

func Set(key string, value interface{}) error {
	return Rdb.Set(Ctx, key, value, 0).Err()
}

func SetEx(key string, value interface{}, expiration time.Duration) error {
	return Rdb.SetEx(Ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return Rdb.Get(Ctx, key).Result()
}

func Del(keys ...string) error {
	return Rdb.Del(Ctx, keys...).Err()
}

func Exists(key string) (bool, error) {
	result, err := Rdb.Exists(Ctx, key).Result()
	return result > 0, err
}

func HSet(key string, values ...interface{}) error {
	return Rdb.HSet(Ctx, key, values...).Err()
}

func HGet(key, field string) (string, error) {
	return Rdb.HGet(Ctx, key, field).Result()
}

func HGetAll(key string) (map[string]string, error) {
	return Rdb.HGetAll(Ctx, key).Result()
}

func Close() error {
	if Rdb != nil {
		return Rdb.Close()
	}
	return nil
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
