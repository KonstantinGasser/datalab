package permissions

import (
	"context"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.user-authentication/config"
	"github.com/KonstantinGasser/datalab/service.user-authentication/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Permissions represents the structure of the Permissions
// in the database
type Permissions struct {
	UserUuid string          `bson:"_id"`
	UserOrgn string          `bson:"user_orgn"`
	Apps     []AppPermission `bson:"apps"`
}

type AppRole int32

// AppPermissions represent one app the user can access
type AppPermission struct {
	AppUuid string  `bson:"uuid" json:"app_uuid"`
	Role    AppRole `bson:"role" json:"role"`
}

// UpdateAppAccess adds new app permissions to the user's permission document
func UpdateAppAccess(ctx context.Context, repo repo.Repo, userUuid string, permssions AppPermission) (*Permissions, error) {
	filter := bson.M{"_id": userUuid}
	query := bson.D{
		{
			Key: "$addToSet",
			Value: bson.M{
				"apps": permssions,
			},
		},
	}
	_, err := repo.UpdateOne(ctx, config.UserAuthDB, config.UserPermissionColl, filter, query, false)
	if err != nil {
		return nil, err
	}

	var newPermissions Permissions
	err = repo.FindOne(ctx, config.UserAuthDB, config.UserPermissionColl, bson.M{"_id": userUuid}, &newPermissions)
	if err != nil {
		return nil, err
	}

	return &newPermissions, nil
}

// GetAppAccess loops up all apps a given user is allowed to access
func GetAppAccess(ctx context.Context, repo repo.Repo, userUuid string) ([]*common.AppPermission, error) {
	filter := bson.M{"_id": userUuid}

	var storedPermissions Permissions
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
