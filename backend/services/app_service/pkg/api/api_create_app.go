package api

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/utils"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
func (srv AppService) CreateApp(ctx context.Context, request *appSrv.CreateAppRequest) (*appSrv.CreateAppResponse, error) {
	// add tracingID from request to context
	ctx = utils.AddValCtx(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.CreateApp] received create app request\n", utils.StringValueCtx(ctx, "tracingID"))

	status, err := srv.app.CreateApp(ctx, srv.mongoC, request)
	if err != nil {
		logrus.Errorf("<%v>[appService.CreateApp] could not create app: %v\n", utils.StringValueCtx(ctx, "tracingID"), err)
		return &appSrv.CreateAppResponse{
			StatusCode: int32(status),
			Msg:        "could not create app",
		}, nil
	}

	return &appSrv.CreateAppResponse{
		StatusCode: int32(status),
		Msg:        "app has been created",
	}, nil
}
