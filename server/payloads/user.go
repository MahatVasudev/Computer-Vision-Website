package payloads

type CreateUser struct {
	Email     string `json:"email"     validate:"required,email"`
	Password  string `json:"password"  validate:"required"`
	Handshake string `json:"handshake" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SetupUser struct {
	Username      string  `json:"username"           validate:"required"`
	FirstName     string  `json:"firstname"          validate:"required"`
	LastName      *string `json:"lastname,omitempty"`
	Gender        string  `json:"gender"             validate:"required"`
	DarkMode      *int    `json:"darkmode"           validate:"required"`
	PreferedColor string  `json:"preferedcolor"      validate:"required"`
	BirthYear     int     `json:"birthyear"          validate:"required"`
}
