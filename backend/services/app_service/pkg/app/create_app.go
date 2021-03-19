package app

import (
	"context"
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"github.com/KonstantinGasser/clickstream/utils/unique"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateApp some docs
func (app app) CreateApp(ctx context.Context, mongo storage.Storage, req *appSrv.CreateAppRequest) (int, string, error) {

	// duplicate names may exists in the system but owners can only hold unique app names
	result, err := mongo.FindOne(ctx, dbName, appCollection, bson.M{"appName": req.GetName(), "ownerUUID": req.GetOwnerUuid()})
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	// if result map is not empty app name already taken
	if len(result) != 0 {
		return http.StatusBadRequest, "", errors.New("duplicated app names are not possible")
	}
	// insert app in db with defaults
	uuid, err := unique.UUID()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	// by default app owner must be in member list
	appMember := append(req.GetMember(), req.GetOwnerUuid())
	data, err := bson.Marshal(AppItem{
		UUID:        uuid,
		AppName:     req.GetName(),
		Description: req.GetDescription(),
		OwnerUUID:   req.GetOwnerUuid(),
		OrgnDomain:  req.GetOrganization(),
		Member:      appMember,
		Settings:    req.GetSettings(),
		AppToken:    "",
	})
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	if err := mongo.InsertOne(ctx, dbName, appCollection, data); err != nil {
		return http.StatusInternalServerError, "", err
	}

	return http.StatusOK, uuid, nil
}
