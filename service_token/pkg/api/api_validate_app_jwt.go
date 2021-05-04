package api

import (
	"context"
	"net/http"

	tokenSrv "github.com/KonstantinGasser/datalab/protobuf/token_service"
	"github.com/KonstantinGasser/datalab/service_token/pkg/jwts"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// ValidateJWT validates authenticated users JWT. Token can be invalid if they have been changes or if they have expired
func (srv TokenServer) VerifyAppToken(ctx context.Context, request *tokenSrv.VerifyAppTokenRequest) (*tokenSrv.VerifyAppTokenResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[tokenService.VerifyUserToken] received request\n", ctx_value.GetString(ctx, "tracingID"))

	appInfo, err := jwts.GetJWTClaims(ctx, request.GetToken(), jwts.SecretApp)
	if err != nil {
		return &tokenSrv.VerifyAppTokenResponse{StatusCode: http.StatusInternalServerError, Msg: "could not get claims"}, nil
	}
	return &tokenSrv.VerifyAppTokenResponse{
		StatusCode: http.StatusOK,
		Msg:        "claims found",
		AppClaims: &tokenSrv.AppClaims{
			AppUuid:        appInfo["sub"].(string),
			OrgnAndAppHash: appInfo["hash"].(string),
			Domain:         appInfo["origin"].(string),
		},
	}, nil
}
