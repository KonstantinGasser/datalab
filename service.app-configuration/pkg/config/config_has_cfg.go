package config

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_config/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_config/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func (cfg config) HasAtLeastOne(ctx context.Context, storage storage.Storage, cfgUUID string) (bool, errors.ErrApi) {

	var cfgItem ConfigItem
	err := storage.FindOne(ctx, cfgDatabase, cfgCollection, bson.M{"_id": cfgUUID}, &cfgItem)
	if err != nil {
		return false, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "could not look up configs",
			Err:    err,
		}
	}

	if len(cfgItem.Funnel) > 0 || len(cfgItem.Campaign) > 0 || len(cfgItem.BtnTime) > 0 {
		return true, nil
	}
	return false, nil
}
