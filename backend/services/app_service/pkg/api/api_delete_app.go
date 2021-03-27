package api

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// DeleteApp is the grpc endpoint to delete an app based on a app uuid
func (srv AppService) DeleteApp(ctx context.Context, request *appSrv.DeleteAppRequest) (*appSrv.DeleteAppResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.DeleteApp] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, err := srv.app.DeleteApp(ctx, srv.storage, request.GetAppUuid())
	if err != nil {
		logrus.Errorf("<%v>[appService.DeleteApp] could not delete app: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.DeleteAppResponse{StatusCode: int32(status), Msg: err.Error()}, nil
	}
	return &appSrv.DeleteAppResponse{
		StatusCode: int32(status),
		Msg:        "app has been deleted",
	}, nil
}
