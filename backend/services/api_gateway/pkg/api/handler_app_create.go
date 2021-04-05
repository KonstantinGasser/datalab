package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// HandlerAppCreate is the api endpoint if a logged in user wants to create a new application
// Involved services:
// - App-Service
func (api API) HandlerAppCreate(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppCreate] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var payload struct {
		Name        string   `json:"app_name"`
		Description string   `json:"app_description"`
		Member      []string `json:"app_member"`
		Settings    []string `json:"app_settings"`
	}
	if err := api.decode(r.Body, &payload); err != nil {
		logrus.Errorf("<%v>[api.HandlerAppCreate] could not decode r.Body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
	}
	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		logrus.Errorf("<%v>[api.HandlerAppCreate] could not get user-claims, they are nil\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New("could not find user claims"), http.StatusForbidden)
		return
	}
	// invoke grpc call to user-service to create requested app
	respApp, err := api.AppClient.Create(r.Context(), &appSrv.CreateRequest{
		Tracing_ID:   ctx_value.GetString(r.Context(), "tracingID"),
		OwnerUuid:    user.GetUuid(),
		Name:         payload.Name,
		Organization: user.GetOrgnDomain(),
		Description:  payload.Description,
		Member:       payload.Member,
		Settings:     payload.Settings,
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppCreate] could not execute grpc.CreateApp: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, err, http.StatusInternalServerError)
		return
	}

	api.onScucessJSON(w, map[string]string{
		"app_uuid": respApp.GetAppUuid(),
		"msg":      respApp.GetMsg(),
	}, int(respApp.GetStatusCode()))
}
