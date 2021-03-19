package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// DataAppDelete represents the data which is required
// in order to delete an app
type DataAppDelete struct {
	AppUUID string `json:"app_uuid"`
}

// HandlerAppDelete is the entry-point to delete an app of a given user in the system
// Involved services:
// - App-Service
func (api API) HandlerAppDelete(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppDelete] received delete app request\n", ctx_value.GetString(r.Context(), "tracingID"))

	var payload DataAppDelete
	if err := api.decode(r.Body, &payload); err != nil {
		logrus.Errorf("<%v>[api.HandlerDeleteApp] could not decode r.Body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not decode passed JSON"), http.StatusBadRequest)
		return
	}
	// invoke grpc call to user-service to delete app
	resp, err := api.AppServiceClient.DeleteApp(r.Context(), &appSrv.DeleteAppRequest{
		AppUuid: payload.AppUUID,
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerDeleteApp] could execute grpc.DeleteApp %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("failed to delete app"), http.StatusInternalServerError)
		return
	}
	api.onScucessJSON(w, map[string]interface{}{}, int(resp.GetStatusCode()))
}
