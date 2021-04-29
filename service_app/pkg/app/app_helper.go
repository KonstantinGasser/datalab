package app

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"github.com/KonstantinGasser/datalab/utils/hash"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// hasToken looks up if an App already as an generated Token
func hasToken(ctx context.Context, storage storage.Storage, appUUID string) error {
	query := bson.D{
		{
			Key:   "_id",
			Value: appUUID,
		},
	}
	var data struct {
		Token string `bson:"app_token"`
	}
	if err := storage.FindOne(ctx, appDatabase, appCollection, query, &data); err != nil {
		return err
	}
	if data.Token != "" {
		return fmt.Errorf("app token already exists")
	}
	return nil
}

// matchAppHash verifies that the request with domain name and app name matches with the database records
// and that the request caller is the owner of the app
func matchAppHash(ctx context.Context, storage storage.Storage, appUUID, callerUUID, domainAndName string) error {
	query := bson.M{"_id": appUUID, "owner_uuid": callerUUID}

	var appData bson.M
	if err := storage.FindOne(ctx, appDatabase, appCollection, query, &appData); err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
		return err
	}

	if _, ok := appData["orgn_and_app_hash"].(string); !ok {
		return fmt.Errorf("could not verify request to create app token")
	}

	requestHash := hash.Sha256([]byte(domainAndName)).String()
	if appData["orgn_and_app_hash"] != requestHash {
		return nil
	}
	return nil
}
