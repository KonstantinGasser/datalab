package api

import (
	"log"
	"net/http"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/api_gateway/pkg/grpcC"
)

const (
	AccessControlAllowOrigin  = "Access-Control-Allow-Origin"
	AccessControlAllowMethods = "Access-Control-Allow-Methods"
)

// API represents the handler functions and middleware
type API struct {
	// Config holds information about
	// how to handle CORS
	cors CORSConfig
	// Route is a mapper between requested URL and handler
	// allows to add middleware in a nice chained way
	route func(path string, h http.HandlerFunc)
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
			http.HandleFunc(path, h)
		},
		// *** Client Dependencies ***
		UserSrvClient: grpcC.NewUserServiceClient(":8001"),
	}
}

// SetUp maps the API routes to the service and specifies the required middleware
func (api API) SetUp() {
	log.Printf("Adding routes to API-Service...\n")

	// ------ ROUTES ------
	api.route("/login", api.WithCors(
		api.HandlerLogin),
	)
}
