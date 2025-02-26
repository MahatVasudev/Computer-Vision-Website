package config

import "fmt"

// PQBConfig : This is the template that contains all the neccessary variables to connect to a postgres session
type PQBConfig struct {
	DBUser     string
	DBHost     string
	DBPassword string
	DBPort     string
	DBSSLMODE  string
	DBName     string
}

var PQEnvs = InitPQDB()

// InitPQDB : It is the default settings if the aurgument in .env file is not provided
func InitPQDB() *PQBConfig {
	return &PQBConfig{
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBSSLMODE:  getEnv("DB_SSLMODE", "disable"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_Name", "on_sight"),
	}
}

// DBConfig : Contains a template for connecting to a postgres session, It simply returs a string and takes in a pointer to PQBConfig struct as an argument
func DBConfig(conf *PQBConfig) string {

	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		conf.DBUser,
		conf.DBPassword,
		conf.DBHost,
		conf.DBName,
		conf.DBSSLMODE,
	)

}
