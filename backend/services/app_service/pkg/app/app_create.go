package app

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/KonstantinGasser/datalabs/backend/services/app_service/pkg/storage"
	"github.com/KonstantinGasser/datalabs/utils/hash"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateApp some docs
func (app app) Create(ctx context.Context, mongo storage.Storage, appItem AppItem) (int, error) {

	// duplicate names may exists in the system but owners can only hold unique app names
	exists, err := mongo.Exists(ctx, appDatabase, appCollection, bson.M{"name": appItem.AppName, "owner_uuid": appItem.OwnerUUID})
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if exists {
		return http.StatusBadRequest, errors.New("duplicated app names are not possible")
	}

	concated := strings.Join([]string{appItem.OrgnDomain, appItem.AppName}, "/")
	orgnAppHash := hash.Sha256([]byte(concated)).String()
	appItem.OrgnAndAppHash = orgnAppHash

	if err := mongo.InsertOne(ctx, appDatabase, appCollection, appItem); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
