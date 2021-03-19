package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// HandlerGetApps is the entry-point to get all apps for a logged in user
func (api API) HandlerGetApps(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerGetApps] received get-apps request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	authedUser := ctx_value.GetAuthedUser(r.Context())
	if authedUser == nil {
		logrus.Errorf("<%v>[api.HandlerGetApps] could not get user-claims, they are nil\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New("could not find user claims"), http.StatusForbidden)
		return
	}
	// invoke grpc to app-service in order for it to create a new application mapped to the authed user
	resp, err := api.AppServiceClient.GetApps(r.Context(), &appSrv.GetAppsRequest{
		UserUuid:   authedUser.GetUuid(),
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerGetApps] could not execute grpc.GetApps: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get apps"), http.StatusInternalServerError)
		return
	}
	if resp.GetStatusCode() != 200 {
		logrus.Errorf("<%v>[api.HandlerGetApps] could not get apps from app-service: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New(resp.GetMsg()), int(resp.GetStatusCode()))
		return
	}
	api.onScucessJSON(w, resp.GetApps(), http.StatusOK)
}
