package login

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/login/jwts"
)

var (
	ErrInvalidToken   = fmt.Errorf("jwt-token is not valid")
	ErrCorruptedToken = fmt.Errorf("jwt-token does no longer confirm with specifications")
)

// IsLoggedIn verifies the authentic of the JWT in order to tell if it is a
// valid token
func IsLoggedIn(ctx context.Context, token string) (*common.TokenClaims, error) {
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
	return &common.TokenClaims{
		Uuid:         uuid,
		Organization: organization,
	}, nil
}
