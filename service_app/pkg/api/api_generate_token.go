package api

import (
	"context"
	"net/http"
	"strings"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
func (srv AppService) GenerateToken(ctx context.Context, in *appSrv.GenerateTokenRequest) (*appSrv.GenerateTokenResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.GenerateToken] received request\n", ctx_value.GetString(ctx, "tracingID"))

	if err := srv.app.HasPermissions(ctx, srv.storage, in.GetCallerUuid(), in.GetAppUuid()); err != nil {
		return &appSrv.GenerateTokenResponse{StatusCode: err.Code(), Msg: err.Info()}, nil
	}

	orgnAndApp := strings.Join([]string{in.GetOrgnDomain(), in.GetAppName()}, "/")
	appToken, err := srv.app.GetTokenClaims(ctx, srv.storage, srv.tokenService, in.GetAppUuid(), in.GetCallerUuid(), orgnAndApp)
	if err != nil {
		logrus.Errorf("<%v>[appService.GenerateToken] could not generate token: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &appSrv.GenerateTokenResponse{StatusCode: err.Code(), Msg: err.Info(), AppToken: ""}, nil
	}
	return &appSrv.GenerateTokenResponse{
		StatusCode: http.StatusOK,
		Msg:        "app token created",
		AppToken:   appToken,
	}, nil
}
