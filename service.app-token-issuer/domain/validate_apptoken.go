package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/jwts"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
)

func (svc apptokenissuer) validateAppToken(ctx context.Context, jwt string) (string, string, errors.ErrApi) {
	claims, err := jwts.Claims(jwt)
	if err != nil {
		return "", "", errors.New(http.StatusInternalServerError, err, "Could not validate App Token")
	}
	return claims.AppUuid, claims.AppOrigin, nil
}
