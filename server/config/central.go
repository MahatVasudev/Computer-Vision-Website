package config

import (
	"log"

	"github.com/joho/godotenv"
)

type CentralConfig struct {
	PORT       string
	PublicHost string
}

var CentralEnvs = initCentralConfig()

func initCentralConfig() CentralConfig {
	var err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return CentralConfig{
		PORT:       getEnv("PORT", ":5000"),
		PublicHost: getEnv("Public_Host", "http://localhost"),
	}

}
