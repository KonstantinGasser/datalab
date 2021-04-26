package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalabs/service_app/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// InitOnNew initializes all required configurations in the database and maps them to the newly created
// app
// func (cfg config) InitOnNew(ctx context.Context, storage storage.Storage, refApp string) error {

// 	if err := storage.InsertOne(ctx, cfgDatabase, funnelCollection, funnel)
// 	return nil
// }

// Upsert updates the app configurations based on the updateFlag.
func (cfg config) UpdateByFlag(ctx context.Context, storage storage.Storage, changes Cfgs, updateFlag string) (int, error) {
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
		return http.StatusBadRequest, fmt.Errorf("invalid update-flag: %s", updateFlag)
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
