package posts

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	sqlDB *sql.DB
	RedDB *redis.Client
}

func NewStore(sqlDB *sql.DB, RedDB *redis.Client) *Store {
	return &Store{
		sqlDB: sqlDB,
		RedDB: RedDB,
	}
}
