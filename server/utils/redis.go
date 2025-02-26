package utils

import (
	"context"

	"github.com/MahatVasudev/Computer-Vision-Website/server/db"
	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
)

func GetRedisSSID(ctx context.Context, ssid string) (string, error) {
	status := db.RedisInstance.Get(ctx, msg.UserSessionKeys(ssid))

	if status.Err() != nil {
		return "", msg.ErrorNotFound
	}

	return status.Val(), nil
}
