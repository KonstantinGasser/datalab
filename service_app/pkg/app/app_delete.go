package app

import (
	"context"
	"errors"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteApp deletes an app based on the provided appUUID
func (app app) Delete(ctx context.Context, storage storage.Storage, appUUID, callerUUID, orgnAndApp string) (int, error) {

	ok, err := app.matchAppHash(ctx, storage, appUUID, callerUUID, orgnAndApp)
	logrus.Warnf("OK:%v, PARAMS:%v\n", ok, orgnAndApp)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if !ok {
		return http.StatusForbidden, errors.New("could not authorize request")
	}
	if err := storage.DeleteOne(ctx, appDatabase, appCollection, bson.M{"_id": appUUID}); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
