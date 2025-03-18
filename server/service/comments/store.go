package comments

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	db    *sql.DB
	redDb *redis.Client
}

func NewStore(db *sql.DB, redDb *redis.Client) *Store {
	return &Store{
		db,
		redDb,
	}
}
