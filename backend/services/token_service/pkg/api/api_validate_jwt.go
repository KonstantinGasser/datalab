package api

import (
	"context"
	"net/http"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
	"github.com/KonstantinGasser/clickstream/backend/services/token_service/pkg/jwts"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// ValidateJWT validates authenticated users JWT. Token can be invalid if they have been changes or if they have expired
func (srv TokenService) ValidateJWT(ctx context.Context, request *tokenSrv.ValidateJWTRequest) (*tokenSrv.ValidateJWTResponse, error) {
	// add tracingID from request to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())

	logrus.Infof("<%v>[tokenService.ValidateJWT] received validation of JWT request\n", ctx_value.GetString(ctx, "tracingID"))
	userInfo, err := jwts.GetJWTClaims(ctx, request.GetJwtToken())
	if err != nil {
		logrus.Errorf("<%v>[tokenService.ValidateJWT] could not get JWT claims: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &tokenSrv.ValidateJWTResponse{
			StatusCode: http.StatusForbidden,
			Msg:        "user is not authenticated",
			IsValid:    false,
			User:       nil,
		}, nil
	}
	return &tokenSrv.ValidateJWTResponse{
		StatusCode: http.StatusOK,
		Msg:        "user is authenticated",
		IsValid:    true,
		User: &tokenSrv.AuthenticatedUser{
			Username:   userInfo["username"].(string),
			Uuid:       userInfo["uuid"].(string),
			OrgnDomain: userInfo["orgnDomain"].(string),
		},
	}, nil
}
