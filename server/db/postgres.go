package db

import (
	"database/sql"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

// PQNewStorage : Used for Initalizing a PostGreSQL Database Server, and Checking its connection
func PQNewStorage(config string) *sql.DB {
	db, err := sql.Open("postgres", config)

	if err != nil {
		panic(err)
	}

	return db
}
