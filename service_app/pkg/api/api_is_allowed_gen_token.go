package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) IsAllowedToGenToken(ctx context.Context, in *appSrv.IsAllowedToGenTokenRequest) (*appSrv.IsAllowedToGenTokenResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.IsAllowedToGenToken] received request\n", ctx_value.GetString(ctx, "tracingID"))

	if err := srv.app.HasPermissions(ctx, srv.storage, in.GetCallerUuid(), in.GetAppUuid()); err != nil {
		logrus.Errorf("<%v>[appService.CreateApp] could not authorize request: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &appSrv.IsAllowedToGenTokenResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			IsAllowed:  false,
		}, nil
	}
	return &appSrv.IsAllowedToGenTokenResponse{
		StatusCode: http.StatusOK,
		Msg:        "Is allowed to operate on App",
		IsAllowed:  true,
	}, nil
}
