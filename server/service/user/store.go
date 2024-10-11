package user

import (
	"context"
	"database/sql"
	"sync"

	"github.com/MahatVasudev/Computer-Vision-Website/server/types"
)

type Store struct {
	db *sql.DB
	mu sync.Mutex
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		mu: sync.Mutex{},
	}
}

func (s Store) CreateUser(ctx context.Context, user *types.User) error {
  s.mu.Lock()

  defer s.mu.Unlock()

  errorChan := make(chan error, 1)

  go func(){
    defer close(errorChan)

    _, err := s.db.ExecContext(ctx, )
  }
}

type responseSingleUser struct {
	user *types.User
	err  error
}

func ScanRowsIntoUsers(rows *sql.Rows) (*types.User, error) {
	resChan := make(chan responseSingleUser)

	go func() {
		defer close(resChan)
		user := new(types.User)
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Avatar,
			&user.CoverPhoto,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			resChan <- responseSingleUser{nil, err}
			return
		}

		resChan <- responseSingleUser{user, nil}
	}()

	select {
	case res := <-resChan:
		return res.user, res.err
	}
}
