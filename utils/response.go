package utils

import (
	"encoding/json"
	"net/http"
)

func RespondValidationError(w http.ResponseWriter, errors map[string]string) {
	RespondWithJSON(w, http.StatusUnprocessableEntity, map[string]interface{}{"errors": errors})
}

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(response)
}
