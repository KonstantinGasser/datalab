package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/utils"
	"github.com/sirupsen/logrus"
)

func (srv AppService) GetApps(ctx context.Context, request *appSrv.GetAppsRequest) (*appSrv.GetAppsResponse, error) {
	// add tracingID from request to context
	ctx = utils.AddValCtx(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.GetApps] received get apps request\n", utils.StringValueCtx(ctx, "tracingID"))

	result, err := srv.app.GetApps(ctx, srv.mongoC, request)
	if err != nil {
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
