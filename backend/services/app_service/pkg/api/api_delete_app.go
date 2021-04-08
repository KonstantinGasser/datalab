package api

import (
	"context"

	appSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/app_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// DeleteApp is the grpc endpoint to delete an app based on a app uuid
<<<<<<< HEAD
func (srv AppService) DeleteApp(ctx context.Context, in *appSrv.DeleteAppRequest) (*appSrv.DeleteAppResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
=======
func (srv AppService) Delete(ctx context.Context, request *appSrv.DeleteRequest) (*appSrv.DeleteResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
>>>>>>> feature_app_token
	logrus.Infof("<%v>[appService.DeleteApp] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, err := srv.app.DeleteApp(ctx, srv.storage, in.GetAppUuid())
	if err != nil {
		logrus.Errorf("<%v>[appService.DeleteApp] could not delete app: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.DeleteResponse{StatusCode: int32(status), Msg: err.Error()}, nil
	}
	return &appSrv.DeleteResponse{
		StatusCode: int32(status),
		Msg:        "app has been deleted",
	}, nil
}
