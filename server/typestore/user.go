package typestore

import "time"

type User struct {
	ID        string    `json:"id"         validate:"required"`
	Username  string    `json:"username"   validate:"required,min=3,max=15"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  *string   `json:"last_name"`
	Email     string    `json:"-"      validate:"required,email"`
	Password  string    `json:"-"          validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDetails struct {
	Avatar        string `json:"avatar"         validate:"required"`
	CoverPhoto    string `json:"cover_photo"    validate:"required"`
	DarkMode      int    `json:"dark_mode"      validate:"required"`
	PreferedColor string `json:"prefered_color" validate:"required"`
	Gender        string `json:"gender"         validate:"required"`
	BirthYear     int    `json:"birth_year"     validate:"required"`
}

type Redis_UserSession struct {
	Id       string    `json:"id"       validate:"required"`
	Username string    `json:"username" validate:"required"`
	IP       string    `json:"ip"       validate:"required"`
	LoggedIn time.Time `json:"loggedin" validate:"required"`
}

type OTP_Redis struct {
	OTP int `json:"otp" validate:"required"`
}
