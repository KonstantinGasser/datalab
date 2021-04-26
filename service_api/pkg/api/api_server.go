package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	tokenSrv "github.com/KonstantinGasser/datalab/protobuf/token_service"
	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/service_api/pkg/grpcC"
	"github.com/sirupsen/logrus"
)

const (
	// AccessControlAllowOrigin describes the allowed request origins
	AccessControlAllowOrigin = "Access-Control-Allow-Origin"
	// AccessControlAllowMethods describes the methods allowed by this API
	AccessControlAllowMethods = "Access-Control-Allow-Methods"
	// AccessControlAllowHeader describes a header -> ???
	AccessControlAllowHeader = "Access-Control-Allow-Headers"
)

// API represents the handler functions and middleware
type API struct {
	// Config holds information about
	// how to handle CORS
	cors CORSConfig
	// Route is a mapper between requested URL and handler
	// allows to add middleware in a nice chained way
	route func(path string, h http.HandlerFunc)
	// onScucessJSON returns a successful response to the client
	// marshaling the passed data allowing to avoid code duplication
	// content-type will always be application/json
	onScucessJSON func(w http.ResponseWriter, data interface{}, status int)
	// onError response to request if an error occurs
	onError func(w http.ResponseWriter, err error, status int)
	// *** Client Dependencies ***
	UserClient     userSrv.UserClient
	TokenSrvClient tokenSrv.TokenClient
	AppClient      appSrv.AppClient
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
		// CORS specification
		cors: cors,
		// route is a custom function allowing to set path and request handler
		// for a given route (similar to the http.HandlerFunc). However having it
		// customs allows to do middleware in a nicer way
		route: func(path string, h http.HandlerFunc) {
			logrus.Infof("[set-up:route] %s\n", path)
			http.HandleFunc(path, h)
		},
		// onScucessJSON returns a marshaled interface{} with a given status code
		// to the client as its response
		onScucessJSON: func(w http.ResponseWriter, data interface{}, status int) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			if err := json.NewEncoder(w).Encode(data); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"status_code": 500, "msg": "an error occurred"}`))
				return
			}
			return
		},
		// onError is a custom function returning a given error back as response.
		// This way code duplication can be avoided
		onError: func(w http.ResponseWriter, err error, status int) {
			http.Error(w, err.Error(), status)
			return
		},
		// *** Client Dependencies ***
		UserClient:     grpcC.NewUserClient(":8001"),
		TokenSrvClient: grpcC.NewTokenClient(":8002"),
		AppClient:      grpcC.NewAppClient(":8003"),
	}
}

// decode is a custom wrapper to decode the request.Body if in JSON.
// Allows to avoid code duplication. Data is decoded into a map[string]interface{}
func (api API) decode(body io.ReadCloser, data interface{}) error {
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
func (api API) encode(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("[api.encode] could not encode data: %v", err)
		return nil, err
	}
	return b, nil
}

// headerAuth is a wrapper to parse the authentication header from a request.
// Function is primarily called from the middleware.WithAuth to get the JWT token.
func (api API) headerAuth(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", errors.New("[api.headerAuth] could not find any Authentication-Header in request")
	}
	return token, nil
}

// getQuery looks for the passed query in the URL returns an empty string
// if not found
func (api API) getQuery(URL *url.URL, query string) string {
	return URL.Query().Get(query)
}

// SetUp sets up all the routes the API has along with all the middleware
// each request required to have
func (api API) SetUp() {
	logrus.Infof("\n*** adding routes to API-Service ***\n")

	// ------ ROUTES ------
	// ***** AUTH: REGISTER & LOGIN *****
	api.route("/api/v1/user/register",
		api.WithTracing(
			api.WithCors(
				api.HandlerUserRegister,
			),
		),
	)
	api.route("/api/v1/user/login",
		api.WithTracing(
			api.WithCors(
				api.HandlerUserLogin,
			),
		),
	)
	// ***********************************
	// ***** VIEW: ACCOUNT + ACTIONS ***************
	api.route("/api/v2/view/account",
		api.WithTracing(
			api.WithCors(
				api.WithAuth(
					api.HandlerAccountDetails,
				),
			),
		),
	)
	api.route("/api/v2/view/account/update",
		api.WithTracing(
			api.WithCors(
				api.WithAuth(
					api.HandlerAccountUpdate,
				),
			),
		),
	)
	// ***********************************
	// ***** VIEW: APP + ACTIONS ***************
	api.route("/api/v2/view/app/details",
		api.WithTracing(
			api.WithCors(
				api.WithAuth(
					api.HandlerAppDetails,
				),
			),
		),
	)
	api.route("/api/v2/view/app/get", // must passed query: uuid=app_uuid
		api.WithTracing(
			api.WithCors(
				api.WithAuth(
					api.HandlerAppGet,
				),
			),
		),
	)
	api.route("/api/v2/view/app/create",
		api.WithTracing(
			api.WithCors(
				api.WithAuth(
					api.HandlerAppCreate,
				),
			),
		),
	)
	api.route("/api/v2/view/app/delete",
		api.WithTracing(
			api.WithCors(
				api.WithAuth(
					api.HandlerAppDelete,
				),
			),
		),
	)
	api.route("/api/v2/view/app/add/member",
		api.WithTracing(
			api.WithCors(
				api.WithAuth(
					api.HandlerAppAddMember,
				),
			),
		),
	)
	api.route("/api/v2/view/app/generate/token",
		api.WithTracing(
			api.WithCors(
				api.WithAuth(
					api.HandlerAppGenerateToken,
				),
			),
		),
	)
	api.route("/api/v2/view/app/update/config",
		api.WithTracing(
			api.WithCors(
				api.WithAuth(
					api.HandlerAppUpdateConfig,
				),
			),
		),
	)
	// api.route("/api/v2/view/app/upload/img",
	// 	api.WithTracing(
	// 		api.WithCors(
	// 			api.WithAuth(
	// 				api.HandlerAppImageUpload,
	// 			),
	// 		),
	// 	),
	// )
	// ***********************************
}
