package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// HandlerAppDelete is the entry-point to delete an app of a given user in the system
// Involved services:
// - App-Service
func (api API) HandlerAppDelete(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppDelete] received request\n", ctx_value.GetString(r.Context(), "tracingID"))

	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		logrus.Errorf("<%v>[api.HandlerAppDelete] could not get user-claims, they are nil\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New("could not find user claims"), http.StatusUnauthorized)
		return
	}
	var payload struct {
		AppUUID string `json:"app_uuid"`
	}
	if err := api.decode(r.Body, &payload); err != nil {
		logrus.Errorf("<%v>[api.HandlerDeleteApp] could not decode r.Body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not decode passed JSON"), http.StatusBadRequest)
		return
	}
	// invoke grpc call to user-service to delete app
	resp, err := api.AppClient.Delete(r.Context(), &appSrv.DeleteRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		AppUuid:    payload.AppUUID,
		CallerUuid: user.GetUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerDeleteApp] could execute grpc.Delete %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("failed to delete app"), http.StatusInternalServerError)
		return
	}
	api.onScucessJSON(w, map[string]interface{}{}, int(resp.GetStatusCode()))
}
