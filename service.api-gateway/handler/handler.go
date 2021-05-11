package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/KonstantinGasser/datalab/service.api-gateway/domain"
	"github.com/sirupsen/logrus"
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
func WithAllowedOrigins(origins ...string) func(*Handler) {
	return func(handler *Handler) {
		handler.allowedOrigins = origins
	}
}

type Handler struct {
	// *** CORS-Configurations ***
	// accessOrigin refers to the AccessControlAllowOrigin Header
	// which can be set in a request
	allowedOrigins []string
	// allowedMethods refers to the allowedControlAllowMethods Header
	// which can be set in a request
	allowedMethods []string
	// allowedHeader refers to the allowedControlAllowHeader Header
	// which can be set in a request
	allowedHeaders []string
	// onSuccessJSON returns a successful response to the client
	// marshaling the passed data allowing to avoid code duplication
	// content-type will always be application/json
	onSuccessJSON func(w http.ResponseWriter, data interface{}, status int)
	// onError response to request if an error occurs
	onError func(w http.ResponseWriter, err string, status int)
	// *** Service Dependencies ***
	domain domain.GatewayLogic
}

func NewHandler(domain domain.GatewayLogic) *Handler {
	return &Handler{
		// *** CORS-Configurations ***
		allowedOrigins: []string{"*"},
		allowedMethods: []string{"GET", "POST", "OPTIONS"},
		allowedHeaders: []string{"*"},
		// onSuccessJSON returns a marshaled interface{} with a given status code
		// to the client as its response
		onSuccessJSON: func(w http.ResponseWriter, data interface{}, status int) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			if err := json.NewEncoder(w).Encode(data); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"status_code": 500, "msg": "an error occurred"}`))
				return
			}
		},
		// onError is a custom function returning a given error back as response.
		// This way code duplication can be avoided
		onError: func(w http.ResponseWriter, err string, status int) {
			http.Error(w, err, status)
		},
		// *** Service Dependencies ***
		domain: domain,
	}
}

// Apply applies an API-CORS configuration on the API instance
func (handler *Handler) Apply(options ...func(*Handler)) {
	for _, option := range options {
		option(handler)
	}
}

// AddRoute maps a route to a given http.HandlerFunc and can proxy middleware before the execution
// of the http.HandlerFunc
func (handler *Handler) Register(route string, h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	logrus.Infof("[handler.Register] %v\n", route)
	var final = h
	for i := len(middleware) - 1; i >= 0; i-- {
		final = middleware[i](final)
	}
	http.HandleFunc(route, final)
}

// WithCors apply the CORS-Configs of the API to a given API-Endpoint
func (handler *Handler) WithCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set CORS-Header to response writer as defined by the api
		w.Header().Set(accessControlAllowOrigin, strings.Join(handler.allowedOrigins, ","))
		w.Header().Set(accessControlAllowMethods, strings.Join(handler.allowedMethods, ","))
		w.Header().Set(accessControlAllowHeader, strings.Join(handler.allowedHeaders, ","))

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// serve next handler
		next(w, r)
	}
}

// decode is a custom wrapper to decode the request.Body if in JSON.
// Allows to avoid code duplication. Data is decoded into a map[string]interface{}
func (handler *Handler) decode(body io.ReadCloser, data interface{}) error {
	if data == nil {
		return fmt.Errorf("passed data can not be nil")
	}
	defer body.Close()

	if err := json.NewDecoder(body).Decode(data); err != nil {
		logrus.Errorf("[api.decode] could not decode r.Body: %v", err)
		return fmt.Errorf("cloud not decode r.Body: %v", err)
	}
	return nil
}

// encode is a custom wrapper to encode any data to a byte slice in order
// for it to be returned to in the response. Allows to avoid code duplication
func (handler *Handler) encode(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("[api.encode] could not encode data: %v", err)
		return nil, err
	}
	return b, nil
}
