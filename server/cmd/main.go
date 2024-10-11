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
)

func main() {
	_, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	db_sql, err := mysqlDataBase()

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db_sql)

	server := api.NewAPIServer(config.CentralEnvs.PORT, db_sql, nil)

	if err = server.Run(); err != nil {
		log.Fatal(err)
	}
}

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

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("MYSQL: Connected SuccessFully.....")
}
