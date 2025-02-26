package utils

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

	err := json.NewDecoder(req.Body).Decode(payload)

	if err != nil {
		return err
	}

	err = Validate.Struct(payload)

	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(writer http.ResponseWriter, status int, value any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)

	return json.NewEncoder(writer).Encode(value)
}

func WriteError(writer http.ResponseWriter, status int, err error) {
	WriteJSON(writer, status, map[string]string{"error": err.Error()})
}
