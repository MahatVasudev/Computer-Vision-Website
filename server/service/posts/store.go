package posts

import (
	"database/sql"
	"sync"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	sqlDB   *sql.DB
	RedDB   *redis.Client
	postLib string
	mu      *sync.Mutex
}

func NewStore(sqlDB *sql.DB, RedDB *redis.Client, posts_dir string) *Store {
	return &Store{
		sqlDB:   sqlDB,
		RedDB:   RedDB,
		postLib: posts_dir,
	}
}
