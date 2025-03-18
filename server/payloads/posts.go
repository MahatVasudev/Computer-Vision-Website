package payloads

type PostPayload struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
}
