package user

import "github.com/MahatVasudev/Computer-Vision-Website/server/typestore"

type temp_email struct {
	Email string `json:"email" validate:"required"`
}

type temp_username struct {
	Username string `json:"username" validate:"required"`
}

type temp_email_otp struct {
	Email    string `json:"email" validate:"required"`
	Otp_code int    `json:"otp"   validate:"required"`
}

type UserWithFollows struct {
	users     typestore.User `json:"user_details"`
	followers int            `json:"follower"`
	following int            `json:"following"`
}
