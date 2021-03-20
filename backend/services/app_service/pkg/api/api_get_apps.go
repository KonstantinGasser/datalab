package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// GetApps serves as the grpc implementation to retrieve all apps created by the logged in user
func (srv AppService) GetApps(ctx context.Context, request *appSrv.GetAppsRequest) (*appSrv.GetAppsResponse, error) {
	// add tracingID from request to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.GetApps] received get apps request\n", ctx_value.GetString(ctx, "tracingID"))

	status, result, err := srv.app.GetApps(ctx, srv.storage, request.GetUserUuid())
	if err != nil {
		logrus.Errorf("<%v>[appService.GetApps] could not get apps: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.GetAppsResponse{StatusCode: int32(status), Msg: "could not get list of apps", Apps: []*appSrv.LightApp{}}, nil
	}

	// convert appList to grpc LightApp slice
	var apps []*appSrv.LightApp = make([]*appSrv.LightApp, len(result))
	for i, item := range result {
		apps[i] = &appSrv.LightApp{Name: item.AppName, Uuid: item.UUID}
	}
	return &appSrv.GetAppsResponse{
		StatusCode: http.StatusOK,
		Msg:        "apps found",
		Apps:       apps,
	}, nil
}
