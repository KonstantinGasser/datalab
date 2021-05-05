package apptoken

import (
	"context"
	"net/http"
	"time"

	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/storage"
	"github.com/KonstantinGasser/datalab/service_config/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (apt apptoken) canBeOverriden(ctx context.Context, db storage.Storage, appUUID string) (bool, errors.ErrApi) {
	var claims TokenClaims
	if err := db.FindOne(ctx, apptokenDB, apptokenColl, bson.M{"_id": appUUID}, &claims); err != nil {
		return false, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not check if override is ok",
			Err:    err,
		}
	}
	currDate := time.Now().Unix()
	if currDate < claims.Exp.Unix() {
		return false, nil
	}
	return true, nil
}
