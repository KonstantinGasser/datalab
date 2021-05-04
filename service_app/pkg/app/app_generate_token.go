package app

import (
	"context"
	"fmt"
	"net/http"

	tokenSrv "github.com/KonstantinGasser/datalab/protobuf/token_service"
	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"go.mongodb.org/mongo-driver/bson"
)

// GetTokenClaims gathers all data required to to generate the app token. It calls out to the TokenServer to issue
// an app token.
func (app app) GetTokenClaims(ctx context.Context, storage storage.Storage, tokenS tokenSrv.TokenClient, appUUID, callerUUID, orgnAndApp string) (string, errors.ErrApi) {
	// pre-conditions:
	//		- orgn name and app name from request must match with db records
	// 		- app must not have a token already
	// 		- app must have at least one configuration set

	if err := matchAppHash(ctx, storage, appUUID, callerUUID, orgnAndApp); err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusBadRequest,
			Err:    fmt.Errorf("matchAppHash: %w", err),
			Msg:    "App-Hash do not match",
		}
	}
	// if err := config.HasConfig(ctx, storage, appUUID); err != nil {
	// 	return "", errors.ErrAPI{
	// 		Status: http.StatusBadRequest,
	// 		Err:    fmt.Errorf("hasConfig: %w", err),
	// 		Msg:    "App needs at least on config set",
	// 	}
	// }
	if err := hasToken(ctx, storage, appUUID); err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusBadRequest,
			Err:    fmt.Errorf("hasToken: %w", err),
			Msg:    "App already has an App-Token",
		}
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
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not issue App-Token",
		}
	}

	resp, err := tokenS.IssueAppToken(ctx, &tokenSrv.IssueAppTokenRequest{
		Tracing_ID:     ctx_value.GetString(ctx, "tracingID"),
		AppUuid:        appData.UUID,
		Origin:         appData.URL,
		OrgnAndAppHash: appData.OrgnAndAppHash,
	})
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not issue App-Token",
		}
	}

	// update database record of app with new app token
	query := bson.D{
		{
			Key:   "$set",
			Value: bson.M{"app_token": resp.GetAppToken()},
		},
	}
	if n, err := storage.UpdateOne(ctx, appDatabase, appCollection, bson.M{"_id": appUUID}, query, false); err != nil || n == 0 {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not issue App-Token",
		}
	}
	return resp.GetAppToken(), nil
}
