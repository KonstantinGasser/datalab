package app

import (
	"context"

	"github.com/KonstantinGasser/datalab/service_app/pkg/config"
	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
)

func (app app) UpdateConfig(ctx context.Context, storage storage.Storage, cfg config.Cfgs, updateFlag string) errors.ErrApi {
	if err := config.UpdateByFlag(ctx, storage, cfg, updateFlag); err != nil {
		return err
	}
	return nil
}
