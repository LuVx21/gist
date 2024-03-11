package db

import (
    "fmt"
    "github.com/redis/go-redis/v9"
    "luvx/config"
)

var RedisClient *redis.Client

func init() {
    fmt.Println("初始化Redis连接...")
    redisConfig := config.AppConfig.Redis
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     redisConfig.Host,
        Username: redisConfig.Username,
        Password: redisConfig.Password,
        DB:       0,
    })
}
