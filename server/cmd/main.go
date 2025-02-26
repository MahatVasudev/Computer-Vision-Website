package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/MahatVasudev/Computer-Vision-Website/server/cmd/api"
	"github.com/MahatVasudev/Computer-Vision-Website/server/config"
	"github.com/MahatVasudev/Computer-Vision-Website/server/db"
	"github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	dbConfig := config.DBConfig(config.InitPQDB())

	pqdb := db.PQNewStorage(dbConfig)

	initNewPQStorage(pqdb)

	redDB := db.RedisCliDefault()

	initNewRedisStorage(ctx, redDB)
	server := api.NewAPIServer(config.CentralEnvs.PORT, pqdb, redDB)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initNewPQStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatalf("Postgres: Connection Failed", err)
	}

	log.Println("Postgres: Connected With Postgres Database...")

}

func initNewRedisStorage(ctx context.Context, redDb *redis.Client) {
	if status := redDb.Ping(ctx); status.Err() != nil {
		log.Fatalf("Redis: Connection Failed", status.Err())
	}

	log.Println("Redis: Connected With Redis Database...")
}

// Deprecated : No Longer Used... Shifted to PostgreSQL
func mysqlDataBase() (*sql.DB, error) {
	db, err := db.NewStorage(mysql.Config{
		DBName:               config.MysqlEnvs.DBName,
		User:                 config.MysqlEnvs.DBUser,
		Passwd:               config.MysqlEnvs.DBPassword,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Deprecated : No Longer Needed as we are shifted to PostgreSQL and PQNewStorage from config package covers this function
func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("MYSQL: Connected SuccessFully.....")
}
