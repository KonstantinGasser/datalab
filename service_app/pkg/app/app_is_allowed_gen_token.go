package app

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
)

func (app app) IsAllowedToGenToken(ctx context.Context, db storage.Storage, callerUUID, appUUID, appHash string) (bool, errors.ErrApi) {
	if err := app.HasPermissions(ctx, db, callerUUID, appUUID); err != nil {
		return false, errors.ErrAPI{
			Status: http.StatusUnauthorized,
			Msg:    "Could not authorize request",
			Err:    err,
		}
	}
	if err := app.matchingAppHash(ctx, db, appUUID, callerUUID, appHash); err != nil {
		return false, errors.ErrAPI{
			Status: http.StatusUnauthorized,
			Msg:    "Could not authorize request",
			Err:    err,
		}
	}
	return true, nil
}
