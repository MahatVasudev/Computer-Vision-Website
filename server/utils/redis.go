package utils

import (
	"context"

	"github.com/MahatVasudev/Computer-Vision-Website/server/db"
	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/redis/go-redis/v9"
)

func ReadEncryptedRedisDataToAStruct(
	ctx context.Context,
	key string,
	payload interface{},
	DecryptionKey []byte,
	redisInstance *redis.Client,
) error {
	if redisInstance == nil {
		redisInstance = db.RedisInstance
	}

	status := redisInstance.Get(ctx, key)

	if status.Err() != nil {
		return status.Err()
	}

	value, err := Decrypt(status.Val(), DecryptionKey)

	if err != nil {
		return err
	}

	if err = MapToStruct(value, &payload); err != nil {
		return err
	}

	return nil
}

func GetRedisSSID(ctx context.Context, ssid string) (string, error) {
	status := db.RedisInstance.Get(ctx, msg.UserSessionKeys(ssid))

	if status.Err() != nil {
		return "", msg.ErrorNotFound
	}

	return status.Val(), nil
}
