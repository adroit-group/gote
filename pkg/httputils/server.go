package httputils

import "net/http"

// ServerHandler is an interface that defines the methods that a http server handler should implement.
// It extends the http.Handler interface, adding a RegisterRoutes method that allows the server to register its routes.
// This is useful for servers that need to register routes dynamically, based on configuration or other factors.
type ServerHandler interface {
	http.Handler
	// RegisterRoutes registers a set of routes on the server, taking a base URL as a parameter.
	RegisterRoutes(baseURL string)
}
