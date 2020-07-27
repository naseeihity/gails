package redis

import (
	"context"
	"gails/app/pkg/config"
	"log"

	"github.com/go-redis/redis/v8"
)

//RedisClient redis客户端
var RedisClient *redis.Client

//Init 初始化redis连接
func Init() {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisCfg.Host,
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis connection err: %v", err)
	}
	log.Printf("Redis connection success: %v", pong)

	RedisClient = rdb
}
