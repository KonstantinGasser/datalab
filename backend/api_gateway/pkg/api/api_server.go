package api

import (
	"net/http"
)

// API represents the handler functions and middleware
type API struct {
	// Config holds information about
	// how to handle CORS
	cfg Config
	// Route is a mapper between requested URL and handler
	// allows to add middleware in a nice chained way
	Route func(path string, h http.HandlerFunc)
}

type Config struct {
	AccessControlAllowOrigin  string
	AccessControlAllowMethods string
}

// NewAPI create and returns a new API struct
func NewAPI(cfg Config) API {
	return API{
		cfg: cfg,
		Route: func(path string, h http.HandlerFunc) {
			http.HandleFunc(path, h)
		},
	}
}
