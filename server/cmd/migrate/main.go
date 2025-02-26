package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MahatVasudev/Computer-Vision-Website/server/config"
	"github.com/MahatVasudev/Computer-Vision-Website/server/db"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Migrations are used for migrating sql queries to prefered driver without touching the pqclient your self
	// You can basically write you sql queries in a migration folder and this will automatically ship
	// those queries to the prefered driver (PostGres in this case)
	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			config.PQEnvs.DBUser,
			config.PQEnvs.DBPassword,
			config.PQEnvs.DBHost,
			config.PQEnvs.DBPort,
			config.PQEnvs.DBName,
			config.PQEnvs.DBSSLMODE,
		),
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

// Deprecated: Using PostGres Now
func MySQLMigrate() {
	_, err := db.NewStorage(mysql.Config{
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

	// driver, err := mysql.WithInstance(db, &mysql.Config{})
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
