package config

import (
	"context"

	"github.com/KonstantinGasser/datalab/service_config/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_config/pkg/storage"
)

const (
	cfgDatabase   = "datalab_config"
	cfgCollection = "config"
)

type Config interface {
	Init(ctx context.Context, storage storage.Storage) (string, errors.ErrApi)
	Get(ctx context.Context, storage storage.Storage, cfgUUID string) (ConfigItem, errors.ErrApi)
	HasAtLeastOne(ctx context.Context, storage storage.Storage, cfgUUID string) (bool, errors.ErrApi)
	Update(ctx context.Context, storage storage.Storage, cfgItem ConfigItem, flag string) errors.ErrApi
}

type config struct{}

type ConfigItem struct {
	UUID     string   `bson:"_id"`
	Funnel   []Stage  `bson:"funnel"`
	Campaign []Record `bson:"campaign"`
	BtnTime  []BtnDef `bson:"btn_time"`
}

type Stage struct {
	ID         int32  `bson:"id"`
	Name       string `bson:"name"`
	Transition string `bson:"transition"`
}
type Record struct {
	ID     int32  `bson:"id"`
	Name   string `bson:"name"`
	Prefix string `bson:"prefix"`
}
type BtnDef struct {
	ID      int32  `bson:"id"`
	Name    string `bson:"name"`
	BtnName string `bson:"btn_name"`
}

func New() Config {
	return &config{}
}
