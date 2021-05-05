package app

import (
	"context"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
)

func (app app) IsAllowedToGenToken(ctx context.Context, callerUUID, appUUID string) (bool, errors.ErrApi) {
	return false, nil
}
