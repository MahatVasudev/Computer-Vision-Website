package writer

import (
	"encoding/json"
	"net/http"

	response_success "github.com/MahatVasudev/Computer-Vision-Website/server/responses/success"
)

func WriteJSON(writer http.ResponseWriter, status int, value any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)

	return json.NewEncoder(writer).Encode(value)
}

func WriteError(writer http.ResponseWriter, status int, err error) {
	WriteJSON(writer, status, map[string]string{"error": err.Error()})
}

func WriteNotFound(writer http.ResponseWriter, err error) {
	WriteError(writer, http.StatusNotFound, err)
}

func WriteOk(w http.ResponseWriter, value any) {
	WriteJSON(w, http.StatusOK, response_success.DataMessage(value))
}

func WriteNotAuthorized(w http.ResponseWriter, err error) {
	WriteError(w, http.StatusUnauthorized, err)
}

func WriteServerError(w http.ResponseWriter, err error) {
	WriteError(w, http.StatusInternalServerError, err)
}

func WriteBadRequest(w http.ResponseWriter, err error) {
	WriteError(w, http.StatusBadRequest, err)
}

func WriteConflictError(writer http.ResponseWriter, err error) {
	WriteError(writer, http.StatusConflict, err)
}
