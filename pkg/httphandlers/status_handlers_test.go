package httphandlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adroit-group/go-template/pkg/version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVersionHandlerFunc(t *testing.T) {
	var versionProvider version.VersionProvider = func() version.Version {
		return version.Version{
			Committish: "abc123",
			BuildDate:  "2023-10-01T00:00:00Z",
		}
	}
	handler := NewVersionHandlerFunc(versionProvider)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var responseBody version.Version
	err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
	require.NoError(t, err)
	require.Equal(t, "abc123", responseBody.Committish)
	require.Equal(t, "2023-10-01T00:00:00Z", responseBody.BuildDate)
}

func TestHealthHandlerFunc(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	HealthHandlerFunc(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var responseBody map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
	require.NoError(t, err)
	require.Equal(t, "ok", responseBody["status"])
}
