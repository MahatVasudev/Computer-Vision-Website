package types

import "time"

type User struct {
	ID         string    `json:"id"          validate:"required"`
	Username   string    `json:"username"    validate:"required,min=3,max=15"`
	FirstName  string    `json:"first_name"  validate:"required"`
	LastName   *string   `json:"last_name"`
	Email      string    `json:"email"       validate:"required,email"`
	Password   string    `json:"-"           validate:"required"`
	Avatar     string    `json:"avatar"      validate:"required"`
	CoverPhoto string    `json:"cover_photo" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserStore interface {
}
