package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

const (
	loginCtxTimeout = time.Second * 2
)

// HandlerUserLogin is the entry-point if a users logins onto the system.
// It calls the user-service to authenticate the users passed
// credentials and on success calls the token-service to issue a new
// JSON-Web-Token holding user information and meta data
// Services involved:
// 	- User-Service
// 	- Token-Service
func (api API) HandlerUserLogin(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerUserLogin] received user-login request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)
	data, err := api.decode(r.Body)
	if err != nil {
		logrus.Errorf("*s*[api.HandlerLogin] could not decode request body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not decode request body"), http.StatusBadRequest)
		return
	}
	// invoke grpc to user-service to check if the passed credentials matches the ones
	// in the database. Response holds the user information for the JWT as well as a bool
	// if authenticated. If not user information will be nil. The status code further tells
	// success of request (403 !authenticated || 500 if service crashed)
	ctx, cancel := context.WithTimeout(r.Context(), loginCtxTimeout)
	defer cancel()
	respUser, err := api.UserSrvClient.AuthUser(ctx, &userSrv.AuthRequest{
		Username:   data["username"].(string),
		Password:   data["password"].(string),
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
	})
	if err != nil || respUser.GetStatusCode() >= http.StatusInternalServerError {
		logrus.Errorf("<%v>[api.HandlerLogin] could not execute grpc.AuthUser: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, fmt.Errorf("could execute grpc.AuthUser: %v", err), http.StatusInternalServerError)
		return
	}
	// return the resp.StatusCode to response if user is not authenticated
	// or the grpc call failed (the returned status code to the user is either 403 or 500)
	if respUser.GetStatusCode() != 200 || !respUser.GetAuthenticated() {
		logrus.Infof("<%v>[api.HandlerLogin] could not authenticate user: code-%d, authed:%v", ctx_value.GetString(r.Context(), "tracingID"), respUser.GetStatusCode(), respUser.GetAuthenticated())
		api.onError(w, errors.New("could not authenticate user"), int(respUser.GetStatusCode()))
		return
	}
	// invoke grpc call to token-service if user is authenticated and issue a JWT for the user
	// token will hold {iat,exp, username, uuid, orgnDomain}
	respToken, err := api.TokenSrvClient.IssueJWT(ctx, &tokenSrv.IssueJWTRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		User: &tokenSrv.AuthenticatedUser{
			Username:   respUser.GetUser().GetUsername(),
			Uuid:       respUser.GetUser().GetUuid(),
			OrgnDomain: respUser.GetUser().GetOrgnDomain(),
		},
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerLogin] could not execute grpc.IssueJWT: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, fmt.Errorf("could execute grpc.IssueJWTa: %v", err), http.StatusInternalServerError)
		return
	}
	// if the token-service request fails return either proper status code
	if respToken.GetStatusCode() != 200 || respToken.GetJwtToken() == "" {
		api.onError(w, errors.New("could not execute authentication request"), int(respToken.GetStatusCode()))
		return
	}
	// return response with JWT
	api.onScucessJSON(w, map[string]interface{}{"token": respToken.GetJwtToken()}, int(respToken.GetStatusCode()))
	return
}
