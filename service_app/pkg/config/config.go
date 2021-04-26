package config

import (
	"context"

	"github.com/KonstantinGasser/datalabs/service_app/pkg/storage"
)

const (
	cfgDatabase        = "datalabs_app"
	funnelCollection   = "funnel"
	campaignCollection = "campaign"
)

type Config interface {
	Upsert(ctx context.Context, storage storage.Storage, cfg Cfgs, updateFlag string) (int, error)
}

type config struct{}

func NewConfig() Config {
	return &config{}
}
