package config

import (
	"context"

	"github.com/KonstantinGasser/datalabs/service_app/pkg/storage"
)

const (
	// MongoDB database name
	cfgDatabase = "datalabs_app"
	// collection to store funnel configs
	appcfgColl = "app"
)

type Config interface {
	UpdateByFlag(ctx context.Context, storage storage.Storage, cfg Cfgs, updateFlag string) (int, error)
	// InitOnNew(ctx context.Context, storage storage.Storage, refApp string) error
}

type config struct{}

func NewConfig() Config {
	return &config{}
}

type Cfgs struct {
	RefApp   string
	Funnel   Funnel   `bson:"funnel"`
	Campaign Campaign `bson:"campaign"`
	BtnTime  BtnTime  `bson:"btn_time"`
}

type Funnel []Stage
type Stage struct {
	ID         int32  `bson:"id" required:"yes"`
	Name       string `bson:"name" required:"yes"`
	Transition string `bson:"transition" required:"yes"`
}

type Campaign []Record

type Record struct {
	ID     int32  `bson:"id" required:"yes"`
	Name   string `bson:"name" required:"yes"`
	Prefix string `bson:"prefix" required:"yes"`
}

type BtnTime []BtnDef

type BtnDef struct {
	ID      int32  `bson:"id" required:"yes"`
	Name    string `bson:"name" required:"yes"`
	BtnName string `bson:"btn_name" required:"yes"`
}
