package api

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) CanGenToken(ctx context.Context, request *appSrv.CanGenTokenRequest) (*appSrv.CanGenTokenResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.CanGenToken] received request\n", ctx_value.GetString(ctx, "tracingID"))

	return nil, nil
}
