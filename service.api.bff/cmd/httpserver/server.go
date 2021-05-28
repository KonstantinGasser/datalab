package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps/collecting"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps/creating"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users/authenticating"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users/fetching"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users/updating"
	"github.com/sirupsen/logrus"
)

func onErr(w http.ResponseWriter, status int32, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	w.Write([]byte(message))
}

func onSuccess(w http.ResponseWriter, status int32, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status_code": 500, "msg": "an error occurred"}`))
		return
	}
}

// WithAllowedOrigins overrides the default setting for
// "Access-Control-Allow-Origin: *" with the given origins
func WithAllowedOrigins(origins ...string) func(*Server) {
	return func(handler *Server) {
		handler.allowedOrigins = origins
	}
}

type Server struct {
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
	// onError response to request if an error occurs
	onErr func(w http.ResponseWriter, status int32, message string)
	// onSuccess returns a successful response to the client
	// marshaling the passed data allowing to avoid code duplication
	// content-type will always be application/json
	onSuccess func(w http.ResponseWriter, status int32, data interface{})
	// *** Server dependiencies ***
	userauthService   authenticating.Service
	userupdateService updating.Service
	userfetchService  fetching.Service
	appcreateService  creating.Service
	appCollectService collecting.Service
}

func NewDefault(
	authService authenticating.Service,
	userupdateService updating.Service,
	userfetchService fetching.Service,
	appcreateService creating.Service,
	appCollectService collecting.Service,
) *Server {
	return &Server{
		// *** CORS-Configurations ***
		allowedOrigins:    []string{"*"},
		allowedMethods:    []string{"GET", "POST", "OPTIONS"},
		allowedHeaders:    []string{"*"},
		onErr:             onErr,
		onSuccess:         onSuccess,
		userauthService:   authService,
		userupdateService: userupdateService,
		userfetchService:  userfetchService,
		appcreateService:  appcreateService,
		appCollectService: appCollectService,
	}
}

func (s Server) Start(host string) error {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	logrus.Infof("[httpserver] starting on: %s\n", host)
	return http.Serve(listener, nil)
}

func (s Server) Register(route string, h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	logrus.Infof("[handler.Register] %v\n", route)
	var final = h
	for i := len(middleware) - 1; i >= 0; i-- {
		final = middleware[i](final)
	}
	http.HandleFunc(route, final)
}

// Apply applies an API-CORS configuration on the API instance
func (s *Server) Apply(options ...func(*Server)) {
	for _, option := range options {
		option(s)
	}
}

func (s Server) decode(body io.ReadCloser, data interface{}) error {
	if data == nil {
		return fmt.Errorf("passed data can not be nil")
	}
	defer body.Close()

	if err := json.NewDecoder(body).Decode(data); err != nil {
		logrus.Errorf("[Server.decode] could not decode r.Body: %v", err)
		return fmt.Errorf("cloud not decode r.Body: %v", err)
	}
	return nil
}
