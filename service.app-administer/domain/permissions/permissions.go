package permissions

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotAuthorized  = fmt.Errorf("caller is not authorized to perform the action")
	ErrNotFound       = fmt.Errorf("could not find app information")
	ErrNoAppHashFound = fmt.Errorf("could not find any app hash for app uuid")
)

func IsOwner(ctx context.Context, repo repo.Repo, callerUuid, appUuid string) error {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.M{"_id": appUuid},
				bson.M{"owner_uuid": callerUuid},
			},
		},
	}
	ok, err := repo.Exists(ctx, config.AppDB, config.AppColl, filter)
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotAuthorized
	}
	return nil
}

// IsCorrectHash checks if the provided app hash matches with the database records
// is used to authorize certain action performed on an app
func IsCorrectHash(ctx context.Context, repo repo.Repo, appUuid, hash string) error {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.M{"_id": appUuid},
				bson.M{"app_hash": hash},
			},
		},
	}
	ok, err := repo.Exists(ctx, config.AppDB, config.AppColl, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotAuthorized
		}
		return err
	}
	// double check if no docs found implies no permissions, right?
	if !ok {
		return ErrNotAuthorized
	}
	return nil
}
