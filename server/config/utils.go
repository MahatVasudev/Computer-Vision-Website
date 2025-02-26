package config

import (
	"os"
	"strconv"
)

func getEnv(env string, fallback string) string {
	if val, ok := os.LookupEnv(env); ok {
		return val
	}
	return fallback

}
func getEnvInt(env string, fallback int) int {
	if val, ok := os.LookupEnv(env); ok {
		if valint, err := strconv.ParseInt(val, 10, 64); err != nil {
			return fallback
		} else {
			return int(valint)
		}

	}
	return fallback

}

func getEnvByte(env string, fallback string) []byte {
	if val, ok := os.LookupEnv(env); ok {
		return []byte(val)
	}

	return []byte(fallback)
}
