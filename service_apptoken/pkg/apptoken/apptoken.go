package apptoken

import (
	"context"
	"time"

	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/storage"
	"github.com/KonstantinGasser/datalab/service_config/pkg/errors"
)

const (
	apptokenDB   = "datalab_apptoken"
	apptokenColl = "apptoken"
	apptokenExp  = (time.Hour * 24) * 7 // valid 7 days
)

type AppToken interface {
	Issue(ctx context.Context, db storage.Storage, claims TokenClaims) (string, errors.ErrApi)
}

type TokenClaims struct {
	AppUuid   string    `bson:"_id" required:"yes"`
	AppHash   string    `bson:"app_hash" required:"yes"`
	AppOrigin string    `bson:"app_origin" required:"yes"`
	AppToken  string    `bson:"app_token"`
	Exp       time.Time `bson:"token_exp"`
}

type apptoken struct{}

func New() AppToken {
	return &apptoken{}
}
