package httputils

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestReadJSONRequest(t *testing.T) {
	type TestData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name        string
		body        string
		contentType string
		expectedErr error
		expected    *TestData
	}{
		{
			name:        "valid JSON request",
			body:        `{"name":"John","age":30}`,
			contentType: "application/json",
			expectedErr: nil,
			expected:    &TestData{Name: "John", Age: 30},
		},
		{
			name:        "invalid content type",
			body:        `{"name":"John","age":30}`,
			contentType: "text/plain",
			expectedErr: ErrInvalidContentType,
			expected:    nil,
		},
		{
			name:        "invalid JSON body",
			body:        `{"name":"John""age":`,
			contentType: "application/json",
			expectedErr: ErrInvalidJSONBody,
			expected:    nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", tc.contentType)

			var data TestData
			err := ReadJSONRequest(req, &data)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, &data)
			}
		})
	}
}

func TestValidateAndReadJSONRequest(t *testing.T) {
	type TestData struct {
		Name string `json:"name" validate:"required"`
		Age  int    `json:"age" validate:"gte=0"`
	}

	testCases := []struct {
		name        string
		body        string
		contentType string
		expectedErr bool
		expected    *TestData
	}{
		{
			name:        "valid JSON request with validation",
			body:        `{"name":"John","age":30}`,
			contentType: "application/json",
			expectedErr: false,
			expected:    &TestData{Name: "John", Age: 30},
		},
		{
			name:        "validation error",
			body:        `{"name":"","age":-5}`,
			contentType: "application/json",
			expectedErr: true,
			expected:    nil,
		},
		{
			name:        "invalid JSON body",
			body:        `{"name":"John","age":}`,
			contentType: "application/json",
			expectedErr: true,
			expected:    nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", tc.contentType)

			var (
				data     TestData
				validate = validator.New()
			)
			err := ValidateAndReadJSONRequest(req, validate, &data)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, &data)
			}
		})
	}
}
