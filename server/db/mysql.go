package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func NewStorage(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		return nil, err
	}

	return db, nil
}
