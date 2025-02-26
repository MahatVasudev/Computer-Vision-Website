package requester

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJson(req *http.Request, payload any) error {
	if req.Body == nil {
		return fmt.Errorf("Couldn't Find the body")
	}

	if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
		return err
	}

	return Validate.Struct(payload)
}
