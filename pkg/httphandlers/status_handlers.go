package httphandlers

import (
	"net/http"

	"github.com/adroit-group/gote/pkg/httputils"
	"github.com/adroit-group/gote/pkg/version"
)

// NewVersionHandlerFunc creates a new HTTP handler function that returns the version information
func NewVersionHandlerFunc(version version.VersionProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httputils.WriteJSONResponse(w, http.StatusOK, version())
	}
}

// HealthHandlerFunc is a simple health check handler that returns a 200 OK response
func HealthHandlerFunc(w http.ResponseWriter, r *http.Request) {
	httputils.WriteJSONResponse(w, http.StatusOK, map[string]string{"status": "ok"})
}
