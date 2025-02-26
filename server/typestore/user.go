package typestore

import "time"

type User struct {
	ID        string    `json:"id"         validate:"required"`
	Username  string    `json:"username"   validate:"required,min=3,max=15"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  *string   `json:"last_name"`
	Email     string    `json:"email"      validate:"required,email"`
	Password  string    `json:"-"          validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDetails struct {
	Avatar        string `json:"avatar"         validate:"required"`
	CoverPhoto    string `json:"cover_photo"    validate:"required"`
	DarkMode      int    `json:"dark_mode"`
	PreferedColor string `json:"prefered_color"`
}

type Redis_UserSession struct {
	Id       string `json:"id" validate:"required"`
	Username string `          validate:"required"`
	IP       string
	LoggedIn time.Time
}

type OTP_Redis struct {
	OTP int `json:"otp" validate:"required"`
}
