package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-administer/errors"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/config"
)

func (svc apptokenissuer) initAppToken(ctx context.Context, appToken AppToken) errors.ErrApi {

	err := svc.dao.InsertOne(ctx, config.TokenDB, config.TokenColl, appToken)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not initialize App Token",
			Err:    err,
		}
	}
	return nil
}
