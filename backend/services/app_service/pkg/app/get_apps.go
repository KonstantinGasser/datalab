package app

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// GetApps collects all apps for a requests owner UUID -> all apps where owner == forUUID will be returned
func (app app) GetApps(ctx context.Context, mongo storage.Storage, forUUID string) (int, []AppItemLight, error) {

	var appList []AppItemLight
	if err := mongo.FindAll(ctx, appDatabase, appCollection, bson.D{{"owner_uuid", forUUID}}, &appList); err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, appList, nil
}
