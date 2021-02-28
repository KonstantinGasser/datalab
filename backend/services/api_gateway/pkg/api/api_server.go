package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/api_gateway/pkg/grpcC"
	"github.com/sirupsen/logrus"
)

const (
	AccessControlAllowOrigin  = "Access-Control-Allow-Origin"
	AccessControlAllowMethods = "Access-Control-Allow-Methods"
	AccessControllAllowHeader = "Access-Control-Allow-Headers"
)

// API represents the handler functions and middleware
type API struct {
	// Config holds information about
	// how to handle CORS
	cors CORSConfig
	// Route is a mapper between requested URL and handler
	// allows to add middleware in a nice chained way
	route func(path string, h http.HandlerFunc)
	// onError response to request if an error occurs
	onError func(w http.ResponseWriter, err error, status int)
	// *** Client Dependencies ***
	UserSrvClient userSrv.UserServiceClient
}

// CORSConfig specifies it CORS policy of the API-Server
type CORSConfig struct {
	Cfgs []struct {
		Header, Value string
	}
}

// New create and returns a new API struct
func New(cors CORSConfig) API {
	return API{
		cors: cors,
		route: func(path string, h http.HandlerFunc) {
			logrus.Infof("[set-up:route] %s\n", path)
			http.HandleFunc(path, h)
		},
		onError: func(w http.ResponseWriter, err error, status int) {
			http.Error(w, err.Error(), status)
			return
		},
		// *** Client Dependencies ***
		UserSrvClient: grpcC.NewUserServiceClient(":8001"),
	}
}

// decode takes an io.ReadCloser (r.Body) and unmarshals the body
func (api API) decode(body io.ReadCloser) (map[string]interface{}, error) {
	defer body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		logrus.Errorf("[api.decode] could not decode r.Body: %v", err)
		return nil, fmt.Errorf("cloud not decode r.Body: %v", err)
	}
	return data, nil
}

// encode takes in a interface and marshals the data for the response
func (api API) encode(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("[api.encode] could not encode data: %v", err)
		return nil, err
	}
	return b, nil
}

// SetUp maps the API routes to the service and specifies the required middleware
func (api API) SetUp() {
	logrus.Infof("Adding routes to API-Service...\n")

	// ------ ROUTES ------
	api.route("/api/v1/user/login", api.WithCors(
		api.HandlerLogin,
	))
	api.route("/api/v1/user/register", api.WithCors(
		api.HandlerRegister,
	))
}
