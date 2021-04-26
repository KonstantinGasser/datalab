package api

import (
	"errors"
	"net/http"

	userSrv "github.com/KonstantinGasser/datalabs/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// HandlerUserUpdate is the entry-point to update the user account of a user
// Involved services:
// - User-Service
func (api API) HandlerAccountUpdate(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerUserUpdate] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var payload struct {
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		OrgnPosition  string `json:"orgn_position"`
		ProfileImgURL string `json:"profile_img_url"`
	}
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
	resp, err := api.UserClient.Update(r.Context(), &userSrv.UpdateRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		CallerUuid: user.GetUuid(),
		User: &userSrv.UpdatableUser{
			FirstName:     payload.FirstName,
			LastName:      payload.LastName,
			OrgnPosition:  payload.OrgnPosition,
			ProfileImgUrl: payload.ProfileImgURL,
		},
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerUpdateUser] could not execute grpc.UpdateUser: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not update user account"), http.StatusInternalServerError)
		return
	}
	api.onScucessJSON(w, map[string]interface{}{"msg": "account updated"}, int(resp.GetStatusCode()))
}
