package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

const (
	createUserTimeout = time.Second * 5
)

// HandlerUserRegister is the entry-point if a users creates a new account.
// It performs sanity checks on the input data and forwards the request
// to the user-service.
// Involved services:
//	- User-Service
func (api API) HandlerUserRegister(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerRegister] received user-register request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)
	data, err := api.decode(r.Body)
	if err != nil {
		api.onError(w, err, http.StatusBadRequest)
		return
	}
	// passed data is not allowed to be empty
	// Todo: create helper func on api to perform checks on N inputs
	if data["username"].(string) == "" || data["password"] == "" || data["orgn_domain"] == "" {
		api.onError(w, fmt.Errorf("missing fields in register request"), http.StatusBadRequest)
		return
	}

	// invoke grpc call to user-service to create the user
	// Response holds only a status-code and a msg (could be an error message)
	ctx, cancel := context.WithTimeout(r.Context(), createUserTimeout)
	defer cancel()
	resp, err := api.UserSrvClient.CreateUser(ctx, &userSrv.CreateUserRequest{
		Username:   data["username"].(string),
		Password:   data["password"].(string),
		OrgnDomain: data["orgn_domain"].(string),
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
	})
	if err != nil {
		api.onError(w, fmt.Errorf("could not execute grpc.CreateUser: %v", err), http.StatusInternalServerError)
		return
	}
	logrus.Infof("<%v>[grpc.CreateUser] status: %d, msg: %s", ctx_value.GetString(r.Context(), "tracingID"), resp.GetStatusCode(), resp.GetMsg())

	// on success write to response
	w.WriteHeader(int(resp.GetStatusCode()))
}
