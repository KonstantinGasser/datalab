package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) Get(ctx context.Context, in *appSrv.GetRequest) (*appSrv.GetResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.GetApp] received request\n", ctx_value.GetString(ctx, "tracingID"))

	app, err := srv.app.Get(ctx, srv.storage, srv.userService, in.GetAppUuid(), in.GetCallerUuid())
	if err != nil {
		logrus.Errorf("<%v>[appService.GetApp] could not GetApp: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &appSrv.GetResponse{StatusCode: err.Code(), Msg: err.Info()}, nil
	}
	return &appSrv.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "app data",
		App:        app,
	}, nil
}
