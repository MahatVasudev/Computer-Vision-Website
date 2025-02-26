package config

import "fmt"

type RedisConfig struct {
	Addrs      string
	Password   string
	DB         int
	Secret_Key []byte
}

var RedisEnv = NewRedisConfigDefault()

func NewRedisConfigDefault() *RedisConfig {
	return &RedisConfig{
		Addrs: fmt.Sprintf(
			"%s:%s",
			getEnv("REDIS_HOST", "localhost"),
			getEnv("REDIS_PORT", "6379"),
		),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       getEnvInt("REDIS_DB", 0),
	}
}
