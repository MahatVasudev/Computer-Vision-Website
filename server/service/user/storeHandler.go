package user

import (
	"context"
	"fmt"

	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
)

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*typestore.User, error) {
	query := `
  SELECT id,username,first_name,last_name,email,password,createdat,updatedat from users where email=$1;
  `

	row, err := s.db.QueryContext(ctx, query, email)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	u := new(typestore.User)

	for row.Next() {
		err := row.Scan(
			&u.ID,
			&u.Username,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
			&u.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
	}

	if u.ID == "" {
		return nil, msg.ErrorNotFound
	}

	return u, nil

}

func (s *Store) CreateUser(ctx context.Context, user typestore.User) error {
	query := `
  INSERT INTO users 
  (
    id, 
    username, 
    first_name, 
    last_name, 
    email, 
    password
  ) VALUES ($1, $2, $3, $4, $5, $6);
  `

	_, err := s.db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password)

	if err != nil {
		fmt.Println(err)
		return msg.ErrorAlreadyExists
	}

	return nil
}

func (s *Store) GetAllUsers(ctx context.Context) ([]typestore.User, error) {
	query := `
  SELECT id, 
    username, 
    first_name, 
    last_name, 
    email, 
    password,
    createdat,
    updatedat
  FROM users
  `

	rows, err := s.db.QueryContext(ctx, query)

	if err != nil {
		return []typestore.User{}, err
	}

	defer rows.Close()

	users_all := []typestore.User{}
	users := new(typestore.User)
	for rows.Next() {
		if err := rows.Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Email,
			&users.Password,
			&users.CreatedAt,
			&users.UpdatedAt); err != nil {
			return []typestore.User{}, err
		}
		users_all = append(users_all, *users)
	}
	if len(users_all) < 1 {
		return users_all, msg.ErrorNotFound
	}
	return users_all, nil

}

func (s *Store) GetUserByUserName(ctx context.Context, username string) (*typestore.User, error) {
	query := `
  SELECT id, username, first_name, last_name, email, password, createdat, updatedat
  FROM users
  WHERE username = $1
  `

	rows, err := s.db.QueryContext(ctx, query, username)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := new(typestore.User)
	for rows.Next() {
		if err := rows.Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Email,
			&users.Password,
			&users.CreatedAt,
			&users.UpdatedAt); err != nil {
			return nil, err
		}
	}
	if users.ID == "" {
		return nil, msg.ErrorNotFound
	}
	return users, nil

}

func (s *Store) GetUserByID(ctx context.Context, userid string) (*typestore.User, error) {
	query := `
  SELECT id, 
    user_name, 
    first_name, 
    last_name, 
    email, 
    password,
    createdat,
    updatedat
  FROM users
  WHERE id = $1
  `

	rows, err := s.db.QueryContext(ctx, query, userid)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := new(typestore.User)
	for rows.Next() {
		if err := rows.Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Email,
			&users.Password,
			&users.CreatedAt,
			&users.UpdatedAt); err != nil {
			return nil, err
		}
	}
	if users.ID == "" {
		return nil, msg.ErrorNotFound
	}
	return users, nil
}
