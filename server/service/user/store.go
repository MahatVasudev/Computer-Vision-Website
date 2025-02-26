package user

import (
	"database/sql"
	"sync"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	db    *sql.DB
	RedDb *redis.Client
	mu    sync.Mutex
}

func NewStore(db *sql.DB, redDB *redis.Client) *Store {
	return &Store{
		db:    db,
		RedDb: redDB,
		mu:    sync.Mutex{},
	}
}
