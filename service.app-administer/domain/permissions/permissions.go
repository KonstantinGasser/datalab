package permissions

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrNotAuthorized = fmt.Errorf("caller is not authorized to perform the action")
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
