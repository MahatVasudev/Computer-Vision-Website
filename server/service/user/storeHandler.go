package user

import (
	"context"
	"fmt"
	"log"

	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
	"github.com/lib/pq"
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

// UpdateUserByID implements store.UserStore.
func (s *Store) UpdateUserByID(ctx context.Context, user typestore.User) error {

	query := `
    UPDATE users SET
      username = COALESCE(NULLIF($2, ''), users.username),
      first_name = COALESCE(NULLIF($3, ''), users.first_name),
      last_name = COALESCE(NULLIF($4, '*'), users.last_name),
      email = COALESCE(NULLIF($5, ''), users.email),
      password = COALESCE(NULLIF($6, ''), users.password)
    WHERE id = $1
  `

	_, err := s.db.ExecContext(ctx, query,
		user.ID,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	)

	if err != nil {
		return err
	}

	return nil

}

// UpdateUserDetailsByID implements store.UserStore.
func (s *Store) UpdateORCreateUserDetailsByID(
	ctx context.Context,
	id string,
	user_details typestore.UserDetails,
) error {

	query := `
    INSERT INTO user_details (userid, avatar, cover_photo, prefered_color, dark_mode, gender, birth_year)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    ON CONFLICT (userid) DO UPDATE
    SET 
    avatar = COALESCE(NULLIF(EXCLUDED.avatar, ''), user_details.avatar),
    cover_photo = COALESCE(NULLIF(EXCLUDED.cover_photo, ''), user_details.cover_photo),
    prefered_color = COALESCE(NULLIF(EXCLUDED.prefered_color, '*'), user_details.prefered_color),
    dark_mode = COALESCE(EXCLUDED.dark_mode, user_details.dark_mode),
    gender = COALESCE(NULLIF(EXCLUDED.gender, ''), user_details.gender),
    birth_year = COALESCE(NULLIF(EXCLUDED.birth_year, 0), user_details.birth_year);
  `

	_, err := s.db.ExecContext(ctx, query,
		id,
		user_details.Avatar,
		user_details.CoverPhoto,
		user_details.PreferedColor,
		user_details.DarkMode,
		user_details.Gender,
		user_details.BirthYear,
	)

	if err != nil {
		return err
	}

	return nil
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

			if errType, ok := err.(*pq.Error); ok {
				log.Println(errType)
			}
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
