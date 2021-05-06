package app

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"github.com/KonstantinGasser/datalab/utils/hash"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteApp deletes an app based on the provided appUUID
func (app app) Delete(ctx context.Context, storage storage.Storage, appUUID, callerUUID, orgnAndApp string) errors.ErrApi {

	appHash := hash.Sha256([]byte(orgnAndApp)).String()
	err := app.matchingAppHash(ctx, storage, appUUID, callerUUID, appHash)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusUnauthorized,
			Err:    err,
			Msg:    "Provided Orgn/App-Name do not match",
		}
	}
	if err := storage.DeleteOne(ctx, appDatabase, appCollection, bson.M{"_id": appUUID}); err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not delete app",
		}
	}
	return nil
}
