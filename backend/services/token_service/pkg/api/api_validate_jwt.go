package api

import (
	"context"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
)

// ValidateJWT validates authenticated users JWT. Token can be invalid if they have been changes or if they have expired
func (srv TokenService) ValidateJWT(context.Context, *tokenSrv.ValidateJWTRequest) (*tokenSrv.ValidateJWTResponse, error) {
	return nil, nil
}
