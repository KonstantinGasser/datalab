package api

import (
	"context"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
)

func (srv TokenService) ValidateJWT(context.Context, *tokenSrv.ValidateJWTRequest) (*tokenSrv.ValidateJWTResponse, error) {
	return nil, nil
}
