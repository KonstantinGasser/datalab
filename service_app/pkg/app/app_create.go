package app

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"github.com/KonstantinGasser/datalab/utils/hash"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateApp some docs
func (app app) Create(ctx context.Context, mongo storage.Storage, appItem AppItem) errors.ErrApi {

	// duplicate names may exists in the system but owners can only hold unique app names
	exists, err := mongo.Exists(ctx, appDatabase, appCollection, bson.M{"name": appItem.AppName, "owner_uuid": appItem.OwnerUUID})
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not create app",
		}
	}
	if exists {
		return errors.ErrAPI{
			Status: http.StatusBadRequest,
			Err:    fmt.Errorf("app name exists already"),
			Msg:    "App names must be unique",
		}
	}

	concated := strings.Join([]string{appItem.OrgnDomain, appItem.AppName}, "/")
	orgnAppHash := hash.Sha256([]byte(concated)).String()
	appItem.OrgnAndAppHash = orgnAppHash

	if err := mongo.InsertOne(ctx, appDatabase, appCollection, appItem); err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not create app",
		}
	}
	return nil
}
