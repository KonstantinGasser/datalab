package api

import (
	"context"
	"net/http"
	"strings"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// DeleteApp is the grpc endpoint to delete an app based on a app uuid
func (srv AppService) Delete(ctx context.Context, in *appSrv.DeleteRequest) (*appSrv.DeleteResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.DeleteApp] received request\n", ctx_value.GetString(ctx, "tracingID"))

	orgnAndApp := strings.Join([]string{in.GetOrgnName(), in.GetAppName()}, "/")
	if err := srv.app.Delete(ctx, srv.storage, in.GetAppUuid(), in.GetCallerUuid(), orgnAndApp); err != nil {
		logrus.Errorf("<%v>[appService.DeleteApp] %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &appSrv.DeleteResponse{StatusCode: err.Code(), Msg: err.Error()}, nil
	}
	return &appSrv.DeleteResponse{
		StatusCode: http.StatusOK,
		Msg:        "app has been deleted",
	}, nil
}
