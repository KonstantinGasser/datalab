package app

import (
	"context"
	"errors"
	"net/http"

	tokenSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/token_service"
	"github.com/KonstantinGasser/datalabs/backend/services/app_service/pkg/storage"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetTokenClaims gathers all data required to to generate the app token. It calls out to the TokenServer to issue
// an app token.
func (app app) GetTokenClaims(ctx context.Context, storage storage.Storage, tokenS tokenSrv.TokenClient, appUUID, callerUUID, orgnAndApp string) (int, string, error) {
	// pre-condition must be true: orgn name and app name from request must match with db records
	ok, err := app.matchAppHash(ctx, storage, appUUID, callerUUID, orgnAndApp)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	if !ok {
		return http.StatusForbidden, "", errors.New("could not authorize request")
	}

	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.M{"_id": appUUID},
				bson.M{"owner_uuid": callerUUID},
			},
		},
	}
	var appData AppItem
	if err := storage.FindOne(ctx, appDatabase, appCollection, filter, &appData); err != nil {
		if err == mongo.ErrNoDocuments {
			return http.StatusForbidden, "", errors.New("could not authorize access")
		}
		return http.StatusInternalServerError, "", err
	}

	resp, err := tokenS.IssueAppToken(ctx, &tokenSrv.IssueAppTokenRequest{
		Tracing_ID:     ctx_value.GetString(ctx, "tracingID"),
		AppUuid:        appData.UUID,
		OrgnAndAppHash: appData.OrgnAndAppHash,
	})
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	// update database record of app with new app token
	query := bson.D{
		{
			Key:   "$set",
			Value: bson.M{"app_token": resp.GetAppToken()},
		},
	}
	if n, err := storage.UpdateOne(ctx, appDatabase, appCollection, bson.M{"_id": appUUID}, query); err != nil || n == 0 {
		return http.StatusInternalServerError, "", errors.New("could not update token in db record")
	}
	return http.StatusOK, resp.GetAppToken(), nil
}
