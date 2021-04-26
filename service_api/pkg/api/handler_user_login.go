package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	tokenSrv "github.com/KonstantinGasser/datalabs/protobuf/token_service"
	userSrv "github.com/KonstantinGasser/datalabs/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
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
	logrus.Infof("<%v>[api.HandlerUserLogin] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := api.decode(r.Body, &payload); err != nil {
		logrus.Errorf("<%v>[api.HandlerLogin] could not decode request body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not decode request body"), http.StatusBadRequest)
		return
	}
	// invoke grpc to user-service to authenticate user
	respUser, err := api.UserClient.Authenticate(r.Context(), &userSrv.AuthenticateRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		Username:   payload.Username,
		Password:   payload.Password,
	})
	if err != nil || respUser.GetStatusCode() >= http.StatusInternalServerError {
		logrus.Errorf("<%v>[api.HandlerLogin] could not execute grpc.AuthUser: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, fmt.Errorf("could execute grpc.AuthUser: %v", err), http.StatusInternalServerError)
		return
	}
	// return the resp.StatusCode to the response if user is not authenticated
	// or the grpc call failed (the returned status code to the user is either 401 or 500)
	if respUser.GetStatusCode() != 200 || respUser.GetUserClaims() == nil {
		logrus.Infof("<%v>[api.HandlerLogin] could not authenticate user: code-%d\n", ctx_value.GetString(r.Context(), "tracingID"), respUser.GetStatusCode())
		api.onError(w, errors.New("could not authenticate user"), int(respUser.GetStatusCode()))
		return
	}

	// invoke grpc to token-service to issue JWT
	respToken, err := api.TokenSrvClient.IssueUserToken(r.Context(), &tokenSrv.IssueUserTokenRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		Claim: &tokenSrv.UserClaim{
			Uuid:       respUser.GetUserClaims().GetUuid(),
			OrgnDomain: respUser.GetUserClaims().GetOrgnDomain(),
		},
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerLogin] could not execute grpc.IssueJWT: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, fmt.Errorf("could execute grpc.IssueJWTa: %v", err), http.StatusInternalServerError)
		return
	}
	// if the token-service request fails return either proper status code
	if respToken.GetStatusCode() != 200 || respToken.GetUserToken() == "" {
		api.onError(w, errors.New("could not execute authentication request"), int(respToken.GetStatusCode()))
		return
	}
	// return response with JWT
	api.onScucessJSON(w, map[string]interface{}{"token": respToken.GetUserToken()}, int(respToken.GetStatusCode()))
	return
}
