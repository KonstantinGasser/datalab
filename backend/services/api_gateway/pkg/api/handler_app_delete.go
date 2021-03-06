package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (api API) HandlerAppDelete(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppDelete] received delete app request\n", ctx_value.GetString(r.Context(), "tracingID"))

	payload, err := api.decode(r.Body)
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerDeleteApp] could not decode r.Body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not decode passed JSON"), http.StatusBadRequest)
		return
	}

	resp, err := api.AppServiceClient.DeleteApp(r.Context(), &appSrv.DeleteAppRequest{
		AppUuid: payload["app_uuid"].(string),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerDeleteApp] could execute grpc.DeleteApp %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("failed to delete app"), http.StatusInternalServerError)
		return
	}
	api.onScucessJSON(w, map[string]interface{}{}, int(resp.GetStatusCode()))
}
