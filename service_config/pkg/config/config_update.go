package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_config/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_config/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func (cfg config) Update(ctx context.Context, storage storage.Storage, cfgItem ConfigItem, flag string) errors.ErrApi {
	var err error
	switch flag {
	case "funnel":
		_, err = storage.UpdateOne(ctx, cfgDatabase, cfgCollection, bson.M{"_id": cfgItem.UUID}, bson.D{
			{
				Key:   "$set",
				Value: bson.M{"funnel": cfgItem.Funnel},
			},
		}, false)
	case "campaign":
		_, err = storage.UpdateOne(ctx, cfgDatabase, cfgCollection, bson.M{"_id": cfgItem.UUID}, bson.D{
			{
				Key:   "$set",
				Value: bson.M{"campaign": cfgItem.Campaign},
			},
		}, false)
	case "btnTime":
		_, err = storage.UpdateOne(ctx, cfgDatabase, cfgCollection, bson.M{"_id": cfgItem.UUID}, bson.D{
			{
				Key:   "$set",
				Value: bson.M{"btn_time": cfgItem.BtnTime},
			},
		}, false)
	default:
		return errors.ErrAPI{
			Status: http.StatusBadRequest,
			Err:    fmt.Errorf("invalid update-flag: %s", flag),
			Msg:    "Provided config flag is invalid",
		}
	}
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not update config(s)",
		}
	}
	return nil
}
