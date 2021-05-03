package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/services/user"
	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/storage"
)

const (
	// accessControlAllowOrigin describes the allowed request origins
	accessControlAllowOrigin = "Access-Control-Allow-Origin"
	// accessControlAllowMethods describes the methods allowed by this API
	accessControlAllowMethods = "Access-Control-Allow-Methods"
	// accessControlAllowHeader describes a header -> ???
	accessControlAllowHeader = "Access-Control-Allow-Headers"
)

// WithAllowedOrigins overrides the default setting for
// "Access-Control-Allow-Origin: *" with the given origins
func WithAllowedOrigins(origins ...string) func(*API) {
	return func(api *API) {
		api.allowedOrigins = origins
		log.Printf("API-Config: %v\n", api.allowedOrigins)
	}
}

// API is the struct which defines the behaviors of
// the web-server. Including routing, middleware and helper funcs
type API struct {
	onError   func(w http.ResponseWriter, code int, err error)
	onSuccess func(w http.ResponseWriter, code int, data interface{})

	// *** CORS-Configurations ***
	allowedOrigins []string
	allowedMethods []string
	allowedHeaders []string
	// *** Service Dependencies ***
	userService user.User
	// *** Storage Dependency ***
	storage storage.Storage
}

// New create a new API-Service with default CORS-Configurations such as
// "Access-Control-Allow-Origin: *", "Access-Control-Allow-Methods: GET,POST,OPTIONS" and
// "Access-Control-Allow-Header: *". To override given settings use api.Apply
func New(userService user.User, storage storage.Storage) API {
	return API{
		onError: func(w http.ResponseWriter, code int, err error) {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
		},
		onSuccess: func(w http.ResponseWriter, code int, data interface{}) {
			if err := json.NewEncoder(w).Encode(data); err != nil {
				log.Printf("could not encode data: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
		// *** CORS-Configurations ***
		allowedOrigins: []string{"*"},
		allowedMethods: []string{"GET", "POST", "OPTIONS"},
		allowedHeaders: []string{"*"},
		// *** Service Dependencies ***
		userService: userService,
		// *** Storage Dependency ***
		storage: storage,
	}
}

// Apply applies configurations to the API-Service
// overriding the default values
func (api *API) Apply(options ...func(*API)) {
	for _, option := range options {
		option(api)
	}
}

// decode takes a io.Reader (can be request body) and decodes the data
// into the passed data interface{}. "data" must be a pointer to the type else
// operation will fail
func (api *API) decode(body io.Reader, data interface{}) error {
	if err := json.NewDecoder(body).Decode(data); err != nil {
		log.Printf("could not decode body: %v", err)
		return err
	}
	return nil
}

// AddRoute maps a route to a given http.HandlerFunc and can proxy middleware before the execution
// of the http.HandlerFunc
func (api *API) AddRoute(route string, h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	defer log.Println("***************************")
	log.Printf("*** Add Route: %s ***\n", route)

	var final = h
	// reverse middleware func so that the first one gets executed fist
	// for i, j := 0, len(middleware)-1; i < j; i, j = i+1, j-1 {
	// 	middleware[i], middleware[j] = middleware[j], middleware[i]
	// }

	for i := len(middleware) - 1; i >= 0; i-- {
		final = middleware[i](final)
	}
	// for _, middle := range middleware {
	// 	final = middle(final)
	// }
	http.HandleFunc(route, final)
}

// WithCors apply the CORS-Configs of the API to a given API-Endpoint
func (api *API) WithCors(next http.HandlerFunc) http.HandlerFunc {
	log.Println("[api.Middleware] applying CORS Headers")
	return func(w http.ResponseWriter, r *http.Request) {
		// set CORS-Header to response writer as defined by the api
		w.Header().Set(accessControlAllowOrigin, strings.Join(api.allowedOrigins, ","))
		w.Header().Set(accessControlAllowMethods, strings.Join(api.allowedMethods, ","))
		w.Header().Set(accessControlAllowHeader, strings.Join(api.allowedHeaders, ","))

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// serve next handler
		next(w, r)
	}
}

// POST checks if the incoming request is a POST request else
// aborts with http.StatusMethodNotAllowed (aka 405)
func (api *API) POST(next http.HandlerFunc) http.HandlerFunc {
	log.Println("[api.Middleware] must be POST check")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// serve next handler
		next(w, r)
	}
}

// GET checks if the incoming request is a GET request else
// aborts with http.StatusMethodNotAllowed (aka 405)
func (api *API) GET(next http.HandlerFunc) http.HandlerFunc {
	log.Println("[api.Middleware] must be GET check")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// serve next handler
		next(w, r)
	}
}
