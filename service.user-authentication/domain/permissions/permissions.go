package permissions

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.user-authentication/config"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-authentication/repo"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateAppAccess(ctx context.Context, repo repo.Repo, userUuid string, permssions types.AppPermission) error {

	query := bson.D{
		{
			Key: "$addToSet",
			Value: bson.M{
				"apps": permssions,
			},
		},
	}
	filter := bson.M{"_id": userUuid}
	_, err := repo.UpdateOne(ctx, config.UserAuthDB, config.UserPermissionColl, filter, query, false)
	if err != nil {
		return err
	}

	return nil
}
