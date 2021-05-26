package domain

import "context"

type Dao interface {
	InsertInitConfig(ctx context.Context, initConfig ConfigInfo) error
	GetById(ctx context.Context, uuid string) (*ConfigInfo, error)
	UpdateByFlag(ctx context.Context, flag string, uuid string, config []interface{}) error
}
