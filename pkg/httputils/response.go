package httputils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// ErrorResponse is a struct that represents an error response.
// It contains an error message and a status code.
type ErrorResponse struct {
	Error  string `json:"error"`  // Error message
	Status int    `json:"status"` // HTTP status code
}

// WriteJSONResponse writes a JSON response to the provided http.ResponseWriter.
//
// The response is encoded as JSON and the provided status code is set on the response.
//
// If the encoding fails, a 500 Internal Server Error response is written.
func WriteJSONResponse[T any](w http.ResponseWriter, status int, response T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}
