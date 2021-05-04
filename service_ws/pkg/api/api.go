package api

import (
	"encoding/json"
	"net/http"
)

// typeKeyCookie is the key to the cookie value in a request
type typeKeyCookie string

// typeKeyDatalabToken is the header key holding the auth-token
type typeKeyDatalabToken string

const (
	// keyCookie is the key to the cookie value in a request
	keyCookie = "x-datalab-cookie"
	// keyDatalabToken is the header key holding the auth-token
	keyDatalabToken = "x-datalab-token"
	// accessControlAllowOrigin refers to the http.Header
	accessControlAllowOrigin = "Access-Control-Allow-Origin"
	// accessControlAllowMethods refers to the http.Header
	accessControlAllowMethods = "Access-Control-Allow-Methods"
	// accessControlAllowHeaders refers to the http.Header
	accessControlAllowHeaders = "Access-Control-Allow-Headers"
	// accessControlAllowCreds refers to the http.Header
	accessControlAllowCreds = "Access-Control-Allow-Credentials"
)

type API struct {
	// CORS config
	allowedOrigins   []string
	allowedMethods   []string
	allowedHeaders   []string
	allowCredentials bool

	onErr     func(w http.ResponseWriter, code int, err string)
	onSuccess func(w http.ResponseWriter, code int, data interface{})

	// *** Dependencies ***
	coordinator string
}

func New(coordinator string) *API {
	return &API{
		// set default CORS config
		allowedOrigins:   []string{"*"},
		allowedMethods:   []string{"GET,POST,OPTIONS"},
		allowedHeaders:   []string{"*"},
		allowCredentials: false,

		onErr: func(w http.ResponseWriter, code int, err string) {
			w.WriteHeader(code)
			w.Write([]byte(err))
		},
		onSuccess: func(w http.ResponseWriter, code int, data interface{}) {
			if err := json.NewEncoder(w).Encode(data); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		},
		coordinator: coordinator,
	}
}

// Route allows to add a new route to the API-Server while also applying passed middleware to the route
func (api *API) Route(route string, h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	// apply set middleware to http route
	var final = h
	for i := len(middleware) - 1; i >= 0; i-- {
		final = middleware[i](final)
	}
	http.HandleFunc(route, final)
}

// Apply overrides default settings of the API
func (api *API) Apply(options ...func(*API)) {
	for _, option := range options {
		option(api)
	}
}

// WithAllowedOrigins apply the given Origins to the
// allowed Origins on the API
func WithAllowedOrgins(origins ...string) func(*API) {
	return func(api *API) {
		api.allowedOrigins = origins
	}
}

// WithAllowedMethods apply the given http.Methods to
// the allowed Methods on the API
func WithAllowedMethods(methods ...string) func(*API) {
	return func(api *API) {
		api.allowedMethods = methods
	}
}

// WithAllowedHeaders apply the given http.Methods to
// the allowed Headers on the API
func WithAllowedHeaders(headers ...string) func(*API) {
	return func(api *API) {
		api.allowedHeaders = headers
	}
}

// WithAllowedCreds set the response Header of
// "Access-Control-Allow-Credentials" to true
func WithAllowedCreds(api *API) {
	api.allowCredentials = true
}
