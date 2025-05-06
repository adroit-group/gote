package httpserver

import (
	"log/slog"
	"net/http"

	"github.com/adroit-group/gote/internal/version"
	"github.com/adroit-group/gote/pkg/httphandlers"
	"github.com/adroit-group/gote/pkg/httputils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type ServerHandler struct {
	mux     *chi.Mux
	valdate *validator.Validate
}

var _ httputils.ServerHandler = (*ServerHandler)(nil)

func (s *ServerHandler) RegisterRoutes(baseURL string) {
	s.mux.Route(baseURL, func(r chi.Router) {
		r.Get("/__version__", httphandlers.NewVersionHandlerFunc(version.GetVersion))
		r.Get("/__health__", httphandlers.HealthHandlerFunc)
	})
	slog.Debug("all routes registered", "baseURL", baseURL)
}

func (s *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// NewServerHandler creates a new ServerHandler.
func NewServerHandler(v *validator.Validate) *ServerHandler {
	return &ServerHandler{
		mux:     chi.NewRouter(),
		valdate: v,
	}
}
