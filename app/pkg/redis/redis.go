package redis

import (
	"context"
	"errors"
	"gails/app/pkg/config"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

//Cache redis cache
type Cache struct {
	Client *redis.Client
	Ctx    context.Context
	Prefix string
}

//RedisClient redis客户端
var RedisClient *redis.Client

//RedisContext redis 上下文
var RedisContext = context.Background()

//Init 初始化redis连接
func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisCfg.Host,
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(RedisContext).Result()
	if err != nil {
		log.Printf("Redis connection err: %v", err)
	}
	log.Printf("Redis connection success: %v", pong)

	RedisClient = rdb
}

//SetCache set redis cache
func (c *Cache) SetCache(key string, value []byte, expire time.Duration) error {
	if RedisClient == nil {
		err := errors.New("Redis not running")
		return err
	}

	cacheKey := c.Prefix + key
	err := c.Client.Set(c.Ctx, cacheKey, value, expire).Err()
	if err != nil {
		return err
	}

	return err
}

//GetCache get redis cache
func (c *Cache) GetCache(key string) ([]byte, error) {
	if RedisClient == nil {
		err := errors.New("Redis not running")
		return nil, err
	}

	cacheKey := c.Prefix + key
	val, err := c.Client.Get(c.Ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	return []byte(val), err
}
