package follows

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	db    *sql.DB
	redDB *redis.Client
}

func NewStore(db *sql.DB, redDB *redis.Client) *Store {
	return &Store{
		db,
		redDB,
	}
}
