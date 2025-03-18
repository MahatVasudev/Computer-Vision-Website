package store

import (
	"context"
	"time"

	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
)

type UserStore interface {

	// SQL Queries
	GetAllUsers(ctx context.Context) ([]typestore.User, error)

	GetUserByID(ctx context.Context, userid string) (*typestore.User, error)

	GetUserByUserName(ctx context.Context, username string) (*typestore.User, error)

	CreateUser(ctx context.Context, user typestore.User) error

	GetUserByEmail(ctx context.Context, email string) (*typestore.User, error)

	UpdateUserByID(ctx context.Context, user typestore.User) error

	UpdateORCreateUserDetailsByID(
		ctx context.Context,
		id string,
		user_details typestore.UserDetails,
	) error

	// Redis Queries

	// UserSession

	CreateUserSession(
		ctx context.Context,
		key string,
		user typestore.Redis_UserSession,
		duration time.Duration,
	) error

	ReadUserKey(ctx context.Context, key string) (*typestore.Redis_UserSession, error)

	// OTP Session
	CreateOTPSession(
		ctx context.Context,
		key string,
		otp int,
		duration time.Duration,
	) error

	DeleteOTPKey(ctx context.Context, key string) error

	ReadOTPKey(ctx context.Context, key string) (*typestore.OTP_Redis, error)
}
