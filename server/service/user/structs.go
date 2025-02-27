package user

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
