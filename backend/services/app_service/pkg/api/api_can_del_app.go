package api

import (
	"context"

	appSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/app_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) CanDelApp(ctx context.Context, request *appSrv.CanDelAppRequest) (*appSrv.CanDelAppResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.CanDelApp] received request\n", ctx_value.GetString(ctx, "tracingID"))

	return nil, nil
}
