package apptoken

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func (apt apptoken) Get(ctx context.Context, db storage.Storage, for_app string) (string, int64, errors.ErrApi) {

	var storedClaims TokenClaims
	if err := db.FindOne(ctx, apptokenDB, apptokenColl, bson.M{"_id": for_app}, &storedClaims); err != nil {
		return "", 0, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not get App-Token Info",
			Err:    err,
		}
	}
	return storedClaims.AppToken, storedClaims.Exp.Unix(), nil
}
