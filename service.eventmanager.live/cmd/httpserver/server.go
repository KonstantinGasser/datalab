package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.eventmanager.live/ports/client"
	"github.com/sirupsen/logrus"
)

// typeKeyCookie is the key to the cookie value in a request
type typeKeyIP string
type typeKeyClaims string
type typeKeyTicket string

// typeKeyDatalabToken is the header key holding the auth-token
type typeKeyDatalabToken string

const (
	// keyCookie is the key to the cookie value in a request
	keyIP = "tracking_ip"
	// keyClaims is the key in the context for the app claims
	keyClaims = "claims"
	// keyTicket is the key in the context holding the web-socket jwt-ticket
	keyTicket = "wsTicket"
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

func onErr(w http.ResponseWriter, status int32, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	w.Write([]byte(message))
}

func onSuccess(w http.ResponseWriter, status int32, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))

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
	allowedHeaders   []string
	allowCredentials bool
	// onError response to request if an error occurs
	onErr func(w http.ResponseWriter, status int32, message string)
	// onSuccess returns a successful response to the client
	// marshaling the passed data allowing to avoid code duplication
	// content-type will always be application/json
	onSuccess func(w http.ResponseWriter, status int32, data interface{})
	// *** Server dependencies ***
	appTokenClient client.ClientAppToken
}

func NewDefault(appTokenClient client.ClientAppToken) *Server {
	srv := &Server{
		// *** CORS-Configurations ***
		allowedOrigins:   []string{"*"},
		allowedMethods:   []string{"GET", "POST", "OPTIONS"},
		allowedHeaders:   []string{"*"},
		allowCredentials: false,
		onErr:            onErr,
		onSuccess:        onSuccess,
		// *** service dependencies ***
		appTokenClient: appTokenClient,
	}
	return srv
}

func (s Server) Start(host string) error {
	// let the stream service listen to incoming events
	// go s.stream.Listen()

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

// WithAllowedOrigins apply the given Origins to the
// allowed Origins on the API
func WithAllowedOrgins(origins ...string) func(*Server) {
	return func(server *Server) {
		server.allowedOrigins = origins
	}
}

// WithAllowedMethods apply the given http.Methods to
// the allowed Methods on the API
func WithAllowedMethods(methods ...string) func(*Server) {
	return func(server *Server) {
		server.allowedMethods = methods
	}
}

// WithAllowedHeaders apply the given http.Methods to
// the allowed Headers on the API
func WithAllowedHeaders(headers ...string) func(*Server) {
	return func(server *Server) {
		server.allowedHeaders = headers
	}
}

// WithAllowedCreds set the response Header of
// "Access-Control-Allow-Credentials" to true
func WithAllowedCreds(server *Server) {
	server.allowCredentials = true
}
