package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

type MysqlConfig struct {
	DBUser     string
	DBPassword string
	DBPort     string
	DBName     string
	DBAddress  string
}

var MysqlEnvs = mysqlinitConfig()

func mysqlinitConfig() MysqlConfig {
	var err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return MysqlConfig{
		DBUser:     getEnv("DB_User", "root"),
		DBPassword: getEnv("DB_Password", "root"),
		DBPort:     getEnv("DB_Port", "3306"),
		DBName:     getEnv("DB_Name", "root"),
		DBAddress:  fmt.Sprint("%s:%s", getEnv("DB_Host", "localhost"), getEnv("DB_Port", "3306")),
	}

}
