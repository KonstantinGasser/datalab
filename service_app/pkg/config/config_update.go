package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateByFlag updates the app configurations based on the updateFlag.
func UpdateByFlag(ctx context.Context, storage storage.Storage, changes Cfgs, updateFlag string) errors.ErrApi {
	var err error
	switch updateFlag {
	case "funnel":
		_, err = storage.UpdateOne(ctx, cfgDatabase, appcfgColl, bson.M{"_id": changes.RefApp}, bson.D{
			{
				Key:   "$set",
				Value: bson.M{"app_config.funnel": changes.Funnel},
			},
		}, false)
	case "campaign":
		_, err = storage.UpdateOne(ctx, cfgDatabase, appcfgColl, bson.M{"_id": changes.RefApp}, bson.D{
			{
				Key:   "$set",
				Value: bson.M{"app_config.campaign": changes.Campaign},
			},
		}, false)
	case "btnTime":
		_, err = storage.UpdateOne(ctx, cfgDatabase, appcfgColl, bson.M{"_id": changes.RefApp}, bson.D{
			{
				Key:   "$set",
				Value: bson.M{"app_config.btn_time": changes.BtnTime},
			},
		}, false)
	default:
		return errors.ErrAPI{
			Status: http.StatusBadRequest,
			Err:    fmt.Errorf("invalid update-flag: %s", updateFlag),
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
