package payloads

type CreateUser struct {
	Email     string `json:"email"     validate:"required,email"`
	Password  string `json:"password"  validate:"required"`
	Handshake string `json:"handshake" validate:"required"`
}

type LoginUser struct {
	Username string `json:"username,omitempty" validate:"required_without=Email"`
	Email    string `json:"email,omitempty"    validate:"required_without=Username"`
	Password string `json:"password"           validate:"required"`
}
