package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// HandlerAppDetails serves the full JSON required by the frontend for the TAB: APP
// actions required by this TAB can be found under HandlerApp* handlerFuns
func (api API) HandlerAppDetails(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppDetails] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	// extract uuid of logged in user
	authedUser := ctx_value.GetAuthedUser(r.Context())
	if authedUser == nil {
		api.onError(w, errors.New("not authenticated"), http.StatusForbidden)
		return
	}

	// fetch all apps mapped to logged in user -> app list in form of: {name, uuid}
	// fetch goes here
	respAppList, err := api.AppServiceClient.GetApps(r.Context(), &appSrv.GetAppsRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		UserUuid:   authedUser.GetUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppDetails] could not execute grpc.GetApps: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get app details information"), http.StatusInternalServerError)
		return
	}
	// if app list is empty or nil return proper data for user
	if respAppList.GetApps() == nil || len(respAppList.GetApps()) == 0 {
		logrus.Infof("<%v>[api.HandlerAppDetails] could not find any apps mapped to authed user\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onScucessJSON(w, map[string]interface{}{
			"app_list":    []struct{}{},
			"app_details": map[string]interface{}{},
		}, http.StatusOK)
		return
	}
	// fetch full app details of first app found in fetch of app list (might be changed to latest updated later)
	// fetch goes here
	respAppDetails, err := api.AppServiceClient.GetByID(r.Context(), &appSrv.GetByIDRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		AppUuid:    respAppList.GetApps()[0].GetUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppDetails] could not execute grpc.GetByID: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get app details information"), http.StatusInternalServerError)
		return
	}
	// build response JSON
	// build goes here
	api.onScucessJSON(w, map[string]interface{}{
		"app_list": respAppList.GetApps(),
		"app_details": map[string]interface{}{
			"app_name":        respAppDetails.GetName(),
			"app_description": respAppDetails.GetDescription(),
			"app_owner":       respAppDetails.GetOwnerUuid(),
			"app_member":      respAppDetails.GetMember(),
			"app_setting":     respAppDetails.GetSettings(),
		},
	}, http.StatusOK)
}
