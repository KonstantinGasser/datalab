package api

import (
	"context"
	"errors"

	tokenSrv "github.com/KonstantinGasser/datalabs/protobuf/token_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// ValidateJWT validates authenticated users JWT. Token can be invalid if they have been changes or if they have expired
func (srv TokenServer) VerifyAppToken(ctx context.Context, request *tokenSrv.VerifyAppTokenRequest) (*tokenSrv.VerifyAppTokenResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[tokenService.VerifyUserToken] received request\n", ctx_value.GetString(ctx, "tracingID"))
	return nil, errors.New("not implemented yet")
}
