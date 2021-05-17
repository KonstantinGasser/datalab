package validate

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/jwts"
)

func Token(ctx context.Context, token string) (string, string, error) {
	claims, err := jwts.Claims(token)
	if err != nil {
		return "", "", nil
	}
	return claims.AppUuid, claims.AppOrigin, nil
}
