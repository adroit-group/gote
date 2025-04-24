package httputils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// ErrInvalidContentType is an error that is returned when the content type of a request is invalid.
var ErrInvalidContentType = errors.New("invalid content type")

// ErrInvalidJSONBody is an error that is returned when the JSON body of a request is invalid.
var ErrInvalidJSONBody = errors.New("invalid JSON body")

// ReadJSONRequest reads a JSON request from the provided http.Request and decodes it into the provided value.
// It returns an error if the content type of the request is not "application/json" or if the decoding fails.
func ReadJSONRequest[T any](r *http.Request, v *T) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return ErrInvalidContentType
	}

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return errors.Join(err, ErrInvalidJSONBody)
	}

	return nil
}

// ValidateAndReadJSONRequest reads a JSON request from the provided http.Request, decodes it into the provided value,
// and validates the value using the provided validator. It returns an error if the content type of the request is not
// "application/json", if the decoding fails, or if the validation fails.
func ValidateAndReadJSONRequest[T any](r *http.Request, v *validator.Validate, t *T) error {
	if err := ReadJSONRequest(r, t); err != nil {
		return err
	}

	return v.Struct(t)
}
