package config

import "os"

func getEnv(env string, fallback string) string {
	if val, ok := os.LookupEnv(env); ok {
		return val
	}
	return fallback

}
