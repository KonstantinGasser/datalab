package api

import (
	"context"
	"errors"
	"net/http"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
	"github.com/KonstantinGasser/clickstream/backend/services/token_service/pkg/jwts"
	"github.com/sirupsen/logrus"
)

// ValidateJWT validates authenticated users JWT. Token can be invalid if they have been changes or if they have expired
func (srv TokenService) ValidateJWT(ctx context.Context, request *tokenSrv.ValidateJWTRequest) (*tokenSrv.ValidateJWTResponse, error) {
	logrus.Info("[tokenService.ValidateJWT] received validation of JWT request\n")
	userInfo, err := jwts.ExtractTokenMetadata(request.GetJwtToken())
	if err != nil {
		return nil, errors.New("could not authenticate user")
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
