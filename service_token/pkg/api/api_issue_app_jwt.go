package api

import (
	"context"
	"net/http"

	tokenSrv "github.com/KonstantinGasser/datalabs/protobuf/token_service"
	"github.com/KonstantinGasser/datalabs/service_token/pkg/jwts"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// IssueJWT issues a new JWT for a authenticated user only
func (srv TokenServer) IssueAppToken(ctx context.Context, in *tokenSrv.IssueAppTokenRequest) (*tokenSrv.IssueAppTokenResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[tokenService.IssueAppToken] received request\n", ctx_value.GetString(ctx, "tracingID"))

	token, err := jwts.IssueApp(ctx, in.GetAppUuid(), in.GetOrgnAndAppHash(), in.GetOrigin())
	if err != nil {
		logrus.Errorf("<%v>[tokenService.IssueAppToken] could not issue JWT for user: %v", ctx_value.GetString(ctx, "tracingID"), err)
		return &tokenSrv.IssueAppTokenResponse{
			StatusCode: http.StatusInternalServerError,
			Msg:        "could not issue JWT for app",
			AppToken:   "",
		}, nil
	}
	return &tokenSrv.IssueAppTokenResponse{
		StatusCode: http.StatusOK,
		Msg:        "JWT for app issued",
		AppToken:   token,
	}, nil
}
