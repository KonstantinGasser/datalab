package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-configuration/errors"
)

func (svc appconfig) initAppConfig(ctx context.Context, initConifg ConfigInfo) errors.ErrApi {
	err := svc.dao.InsertInitConfig(ctx, initConifg)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not initialize App Config")
	}
	return nil
}
