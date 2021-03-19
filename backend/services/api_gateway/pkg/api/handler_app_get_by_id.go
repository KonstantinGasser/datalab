package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (api API) HandlerAppGetByID(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppGetByID] received get app by id request\n", ctx_value.GetString(r.Context(), "tracingID"))

	uuid := api.getQuery(r.URL, "uuid")
	if uuid == "" {
		api.onError(w, errors.New("could not find any app uuid"), http.StatusBadRequest)
		return
	}

	resp, err := api.AppServiceClient.GetByID(r.Context(), &appSrv.GetByIDRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		AppUuid:    uuid,
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppGetByID] could execute grpc.GetByID %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("failed to get app information"), http.StatusInternalServerError)
		return
	}
	if resp.GetStatusCode() != 200 {
		api.onError(w, errors.New("could not fetch app information"), int(resp.GetStatusCode()))
		return
	}
	api.onScucessJSON(w, resp, int(resp.GetStatusCode()))
}
