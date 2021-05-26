package domain

import (
	"context"
	"time"
)

type Dao interface {
	InsertOne(ctx context.Context, db, collection string, query interface{}) error
	GetById(ctx context.Context, uuid string, result interface{}) error
	UpdateAppToken(ctx context.Context, uuid, jwt, origin string, newExp time.Time) error
	HasRWAccess(ctx context.Context, uuid, ownerUuid string) error
}
