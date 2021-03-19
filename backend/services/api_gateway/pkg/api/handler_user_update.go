package api

import (
	"errors"
	"net/http"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

type DataUserUpdate struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	OrgnPosition string `json:"orgn_position"`
}

// HandlerUserUpdate is the entry-point to update the user account of a user
// Involved services:
// - User-Service
func (api API) HandlerUserUpdate(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerUserUpdate] received user update request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var payload DataUserUpdate
	if err := api.decode(r.Body, &payload); err != nil {
		api.onError(w, errors.New("could not decode request body"), http.StatusBadRequest)
		return
	}

	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		logrus.Errorf("<%v>[api.HandlerUpdateUser] could not find authenticated user in context\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New("could not find authenticated user"), http.StatusUnauthorized)
		return
	}
	// call out to user service to update authed user
	resp, err := api.UserSrvClient.UpdateUser(r.Context(), &userSrv.UpdateUserRequest{
		Tracing_ID:   ctx_value.GetString(r.Context(), "tracingID"),
		UUID:         user.Uuid,
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		OrgnPosition: payload.OrgnPosition,
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerUpdateUser] could not execute grpc.UpdateUser: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not update user account"), http.StatusInternalServerError)
		return
	}
	api.onScucessJSON(w, map[string]interface{}{"msg": "account updated"}, int(resp.GetStatusCode()))
}
