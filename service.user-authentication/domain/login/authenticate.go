package login

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/login/jwts"
	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
)

var (
	ErrInvalidToken   = fmt.Errorf("jwt-token is not valid")
	ErrCorruptedToken = fmt.Errorf("jwt-token does no longer confirm with specifications")
)

func IsLoggedIn(ctx context.Context, token string) (*proto.Claims, error) {
	rawClaims, err := jwts.GetJWTClaims(token)
	if err != nil {
		if err == jwts.ErrInvalidJWT {
			return nil, ErrInvalidToken
		}
		return nil, err
	}

	uuid, ok := rawClaims["sub"].(string)
	if !ok {
		return nil, ErrCorruptedToken
	}
	organization, ok := rawClaims["orgn"].(string)
	if !ok {
		return nil, ErrCorruptedToken
	}
	return &proto.Claims{
		Uuid:         uuid,
		Organization: organization,
	}, nil
}
