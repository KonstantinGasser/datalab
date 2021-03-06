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

	result, err := srv.app.GetApps(ctx, srv.mongoC, request)
	if err != nil {
		logrus.Errorf("<%v>[appService.GetApps] could not get apps: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.GetAppsResponse{
			StatusCode: http.StatusInternalServerError,
			Msg:        "could not get list of apps",
			Apps:       []*appSrv.LightApp{},
		}, nil
	}
	return &appSrv.GetAppsResponse{
		StatusCode: http.StatusOK,
		Msg:        "apps found",
		Apps:       result,
	}, nil
}
