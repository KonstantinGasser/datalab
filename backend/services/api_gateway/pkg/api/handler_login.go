package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/sirupsen/logrus"
)

// HandlerLogin forwards the login request to the user and token service
// in oder to check the user's auth and to issue a JWT.
func (api API) HandlerLogin(w http.ResponseWriter, r *http.Request) {

	data, err := api.decode(r.Body)
	if err != nil {
		logrus.Errorf("[api.HandlerLogin] could not decode request body: %v\n", err)
		api.onError(w, errors.New("could not decode request body"), http.StatusBadRequest)
		return
	}
	// // generate password hash
	// hashedPassword, err := utils.HashFromPassword([]byte(data["password"].(string)))
	// if err != nil {
	// 	logrus.Errorf("[api.HandlerLogin] %v\n", err)
	// 	api.onError(w, fmt.Errorf("could not hash password: %v", err), http.StatusInternalServerError)
	// 	return
	// }
	// invoke grpc call
	respUser, err := api.UserSrvClient.AuthUser(context.Background(), &userSrv.AuthRequest{
		Username: data["username"].(string),
		Password: data["password"].(string),
	})
	if err != nil {
		logrus.Errorf("[api.HandlerLogin] could not execute grpc.AuthUser: %v\n", err)
		api.onError(w, fmt.Errorf("could execute grpc.AuthUser: %v", err), http.StatusInternalServerError)
		return
	}
	// return 403 if user service cannot match requested credentials with DB-Result
	if respUser.GetStatusCode() != 200 || !respUser.GetAuthenticated() {
		logrus.Errorf("[api.HandlerLogin] could not authenticate user: code-%d, authed:%v", respUser.GetStatusCode(), respUser.GetAuthenticated())
		api.onError(w, errors.New("could not authenticate user"), http.StatusForbidden)
		return
	}
	// Todo invoke grpc to TokenService to issue a new JWT (on-success)
	respToken, err := api.TokenSrvClient.IssueJWT(context.Background(), &tokenSrv.IssueJWTRequest{
		User: &tokenSrv.AuthenticatedUser{
			Username:   respUser.GetUser().GetUsername(),
			Uuid:       respUser.GetUser().GetUuid(),
			OrgnDomain: respUser.GetUser().GetOrgnDomain(),
		},
	})
	if err != nil {
		logrus.Errorf("[api.HandlerLogin] could not execute grpc.IssueJWT: %v\n", err)
		api.onError(w, fmt.Errorf("could execute grpc.IssueJWTa: %v", err), http.StatusInternalServerError)
		return
	}
	if respToken.GetStatusCode() != 200 || respToken.GetJwtToken() == "" {
		api.onError(w, errors.New("could not execute authentication request"), http.StatusForbidden)
		return
	}
	// return JW-Token
	b, err := api.encode(map[string]interface{}{"token": respToken.GetJwtToken()})
	if err != nil {
		logrus.Errorf("[api.HandlerLogin] could not encode response: %v", err)
		api.onError(w, errors.New("could not encode response"), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(int(respToken.GetStatusCode()))
	w.Write(b)
	return
}
