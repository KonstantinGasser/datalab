package permissions

import (
	"context"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.user-authentication/config"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-authentication/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateAppAccess(ctx context.Context, repo repo.Repo, userUuid string, permssions types.AppPermission) (*types.Permissions, error) {

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
		return nil, err
	}

	var newPermissions types.Permissions
	err = repo.FindOne(ctx, config.UserAuthDB, config.UserPermissionColl, bson.M{"_id": userUuid}, &newPermissions)
	if err != nil {
		return nil, err
	}

	return &newPermissions, nil
}

func GetAppAccess(ctx context.Context, repo repo.Repo, userUuid string) ([]*common.AppPermission, error) {
	filter := bson.M{"_id": userUuid}

	var storedPermissions types.Permissions
	err := repo.FindOne(ctx, config.UserAuthDB, config.UserPermissionColl, filter, &storedPermissions)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []*common.AppPermission{}, nil
		}
		return nil, err
	}
	var permissions = make([]*common.AppPermission, len(storedPermissions.Apps))
	for i, item := range storedPermissions.Apps {
		permissions[i] = &common.AppPermission{AppUuid: item.AppUuid, Role: common.AppRole(item.Role)}
	}
	return permissions, nil
}
