package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong"))
	}
	w.WriteHeader(status)

}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return nil
}

func WriteJSONError(w http.ResponseWriter, status int, message string) {
	type data struct {
		Err string `json:"error"`
	}
	newError := data{Err: message}
	WriteJSON(w, status, newError)
}
