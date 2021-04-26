package app

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalabs/service_app/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// GetApps collects all apps for a requests owner UUID -> all apps where owner == forUUID will be returned
func (app app) GetList(ctx context.Context, mongo storage.Storage, forUUID string) (int, []AppItemLight, error) {

	var appList []AppItemLight
	if err := mongo.FindMany(ctx, appDatabase, appCollection, bson.D{{Key: "owner_uuid", Value: forUUID}}, &appList); err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, appList, nil
}
