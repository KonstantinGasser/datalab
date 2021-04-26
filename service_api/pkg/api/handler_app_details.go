package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// HandlerAppDetails serves the full JSON required by the frontend for the TAB: APP
// actions required by this TAB can be found under HandlerApp* handlerFuns
func (api API) HandlerAppDetails(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppDetails] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	// extract uuid of logged in user
	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		api.onError(w, errors.New("not authenticated"), http.StatusForbidden)
		return
	}

	// fetch all apps mapped to logged in user -> app list in form of: {name, uuid}
	// fetch goes here
	respAppList, err := api.AppClient.GetList(r.Context(), &appSrv.GetListRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		CallerUuid: user.GetUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppDetails] could not execute grpc.GetAppList: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get app details information"), http.StatusInternalServerError)
		return
	}
	// if app list is empty or nil return proper data for user
	if respAppList.GetAppList() == nil || len(respAppList.GetAppList()) == 0 {
		logrus.Infof("<%v>[api.HandlerAppDetails] could not find any apps mapped to caller\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onScucessJSON(w, map[string]interface{}{
			"app_list":    []struct{}{},
			"app_details": map[string]interface{}{},
		}, http.StatusOK)
		return
	}
	// fetch full app details of first app found in fetch of app list (might be changed to latest updated later)
	// fetch goes here
	respAppDetails, err := api.AppClient.Get(r.Context(), &appSrv.GetRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		AppUuid:    respAppList.GetAppList()[0].GetUuid(),
		CallerUuid: user.GetUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppDetails] could not execute grpc.Get: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get app details information"), http.StatusInternalServerError)
		return
	}
	// build response JSON
	// build goes here
	api.onScucessJSON(w, map[string]interface{}{
		"app_list":    respAppList.GetAppList(),
		"app_details": respAppDetails.GetApp(),
	}, http.StatusOK)
}
