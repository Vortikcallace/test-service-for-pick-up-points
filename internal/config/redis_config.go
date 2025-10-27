package config

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func NewConfig(v *viper.Viper) *RedisConfig {
	return &RedisConfig{
		Host:     v.GetString("redis.host"),
		Port:     v.GetInt("redis.port"),
		Password: v.GetString("redis.password"),
		DB:       v.GetInt("redis.db"),
	}
}

func InitRedisClient(cfg *RedisConfig) *redis.Client {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: 10,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis")
	return redisClient
}
