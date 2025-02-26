package db

import (
	"github.com/MahatVasudev/Computer-Vision-Website/server/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cnfg redis.Options) *redis.Client {
	client := redis.NewClient(&cnfg)

	return client
}

func RedisCliDefault() *redis.Client {
	option := redis.Options{
		DB:       config.RedisEnv.DB,
		Addr:     config.RedisEnv.Addrs,
		Password: config.RedisEnv.Password,
	}

	return NewRedisClient(option)
}

var RedisInstance = RedisCliDefault()
