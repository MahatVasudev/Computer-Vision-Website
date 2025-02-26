package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// Deprecated : Now Using PostGres, use PQNewStorage Instead
//
// NewStorage : It is Used for initializing a MySQL Database Server
func NewStorage(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		return nil, err
	}

	return db, nil
}
