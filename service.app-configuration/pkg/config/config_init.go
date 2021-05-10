package config

import (
	"context"
	"log"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_config/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_config/pkg/storage"
	"github.com/KonstantinGasser/datalab/utils/unique"
)

func (cfg config) Init(ctx context.Context, storage storage.Storage) (string, errors.ErrApi) {

	uuid, err := unique.UUID()
	if err != nil {
		return "", errors.ErrAPI{Status: http.StatusInternalServerError, Msg: "Could not init Config", Err: err}
	}
	cfgItem := ConfigItem{UUID: uuid}
	log.Println(cfgItem)
	if err := storage.InsertOne(ctx, cfgDatabase, cfgCollection, cfgItem); err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not init config",
			Err:    err,
		}
	}
	return uuid, nil
}
