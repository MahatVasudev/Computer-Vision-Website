package main

import (
	"log"
	"os"

	"github.com/MahatVasudev/Computer-Vision-Website/server/config"
	"github.com/MahatVasudev/Computer-Vision-Website/server/db"
	mysqlCnfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := db.NewStorage(mysqlCnfg.Config{
		DBName:               config.MysqlEnvs.DBName,
		User:                 config.MysqlEnvs.DBUser,
		Passwd:               config.MysqlEnvs.DBPassword,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]

	if cmd == "up" {
		if err = m.Up(); err != nil {
			log.Fatal(err)
		}
	} else if cmd == "down" {
		if err = m.Down(); err != nil {
			log.Fatal(err)
		}
	}
}
