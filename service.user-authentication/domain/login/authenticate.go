package login

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.user-authentication/jwts"
)

var (
	ErrInvalidToken   = fmt.Errorf("jwt-token is not valid")
	ErrCorruptedToken = fmt.Errorf("jwt-token does no longer confirm with specifications")
)

// IsLoggedIn verifies the authentic of the JWT in order to tell if it is a
// valid token
func IsLoggedIn(ctx context.Context, token string) (*common.UserTokenClaims, error) {
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
	permissions := rawClaims["apps"].([]interface{})
	if !ok {
		return nil, ErrCorruptedToken
	}
	var appPermissions = make([]*common.AppPermission, len(permissions))
	if len(permissions) != 0 {
		for i, item := range permissions {
			tmp := item.(map[string]interface{})
			appPermissions[i] = &common.AppPermission{
				AppUuid: tmp["app_uuid"].(string),
				Role:    common.AppRole(tmp["role"].(float64)),
			}
		}
	}
	return &common.UserTokenClaims{
		Uuid:         uuid,
		Organization: organization,
		Permissions:  &common.UserPermissions{Apps: appPermissions},
	}, nil
}
