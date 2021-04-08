package api

import (
	"context"
	"strings"

	appSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/app_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
func (srv AppService) GenerateToken(ctx context.Context, in *appSrv.GenerateTokenRequest) (*appSrv.GenerateTokenResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.GenerateToken] received request\n", ctx_value.GetString(ctx, "tracingID"))

	orgnAndApp := strings.Join([]string{in.GetOrgnDomain(), in.GetAppName()}, "/")
	status, appToken, err := srv.app.GetTokenClaims(ctx, srv.storage, srv.tokenService, in.GetAppUuid(), in.GetCallerUuid(), orgnAndApp)
	if err != nil {
		logrus.Errorf("<%v>[appService.GenerateToken] could not generate token: %v\n", err, ctx_value.GetString(ctx, "tracingID"))
		return &appSrv.GenerateTokenResponse{StatusCode: int32(status), Msg: "could not generate token", AppToken: ""}, nil
	}
	// if err != nil {
	// 	logrus.Errorf("<%v>[appService.GenerateToken] could not execute grpc.IssueAppToken: %v\n", err, ctx_value.GetString(ctx, "tracingID"))
	// 	return &appSrv.GenerateTokenResponse{StatusCode: http.StatusInternalServerError, Msg: "could not issue app token", AppToken: ""}, nil
	// }
	return &appSrv.GenerateTokenResponse{
		StatusCode: int32(status),
		Msg:        "app token created",
		AppToken:   appToken,
	}, nil
}
