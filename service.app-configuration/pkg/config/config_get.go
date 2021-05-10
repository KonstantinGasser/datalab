package config

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_config/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_config/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func (cfg config) Get(ctx context.Context, storage storage.Storage, cfgUUID string) (ConfigItem, errors.ErrApi) {

	var cfgItem ConfigItem
	if err := storage.FindOne(ctx, cfgDatabase, cfgCollection, bson.M{"_id": cfgUUID}, &cfgItem); err != nil {
		return ConfigItem{}, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not query for config",
			Err:    err,
		}
	}
	return cfgItem, nil
}
