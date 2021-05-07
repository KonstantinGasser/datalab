package apptoken

import (
	"context"
	"time"

	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/storage"
)

const (
	apptokenDB   = "datalab_apptoken"
	apptokenColl = "apptoken"
	apptokenExp  = (time.Hour * 24) * 7 // valid 7 days
)

type AppToken interface {
	Issue(ctx context.Context, db storage.Storage, claims TokenClaims) (*MetaToken, errors.ErrApi)
	Get(ctx context.Context, db storage.Storage, for_app string) (*MetaToken, errors.ErrApi)
}

type TokenClaims struct {
	AppUuid   string    `bson:"_id" required:"yes"`
	AppHash   string    `bson:"app_hash" required:"yes"`
	AppOrigin string    `bson:"app_origin" required:"yes"`
	AppToken  string    `bson:"app_token"`
	Exp       time.Time `bson:"token_exp"`
}

type MetaToken struct {
	Token string
	Exp   int64
}

type apptoken struct{}

func New() AppToken {
	return &apptoken{}
}
