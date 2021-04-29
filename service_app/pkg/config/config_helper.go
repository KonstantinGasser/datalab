package config

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// hasConfig looks up if a given App has at least one configuration set
func HasConfig(ctx context.Context, storage storage.Storage, appUUID string) error {
	query := bson.D{
		{
			Key:   "_id",
			Value: appUUID,
		},
	}

	var data struct {
		Cfg struct {
			Funnel   []Stage  `bson:"funnel"`
			Campaign []Record `bson:"campaign"`
			BtnTime  []BtnDef `bson:"btn_time"`
		} `bson:"app_config"`
	}
	if err := storage.FindOne(ctx, cfgDatabase, appcfgColl, query, &data); err != nil {
		return err
	}
	// verify that app configs are actually set not just empty lists
	if len(data.Cfg.Funnel) == 0 && len(data.Cfg.Campaign) == 0 && len(data.Cfg.BtnTime) == 0 {
		return fmt.Errorf("app has no configs set")
	}
	return nil
}
