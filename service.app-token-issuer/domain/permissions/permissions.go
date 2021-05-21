package permissions

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/config"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotAuthorized  = fmt.Errorf("caller is not authorized to perform the action")
	ErrNotFound       = fmt.Errorf("could not find app information")
	ErrNoAppHashFound = fmt.Errorf("could not find any app hash for app uuid")
	ErrNoAppsFound    = fmt.Errorf("could not find any apps the user is allowed to have access to")
)

func IsOwner(ctx context.Context, repo repo.Repo, callerUuid, appUuid string) error {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.M{"_id": appUuid},
				bson.M{"app_owner": callerUuid},
			},
		},
	}
	ok, err := repo.Exists(ctx, config.TokenDB, config.TokenColl, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotAuthorized
		}
		return err
	}
	if !ok {
		return ErrNotAuthorized
	}
	return nil
}

func CanAccess(ctx context.Context, repo repo.Repo, permissions *common.UserTokenClaims, appUuid string) error {
	fmt.Printf("permissions: %+v\n", permissions)
	if permissions.GetPermissions().GetApps() == nil || len(permissions.GetPermissions().GetApps()) == 0 {
		return ErrNotAuthorized
	}
	var allowedApps []string
	for _, item := range permissions.GetPermissions().GetApps() {
		allowedApps = append(allowedApps, item.GetAppUuid())
	}
	filter := bson.D{
		{
			Key: "_id",
			Value: bson.D{
				{
					Key:   "$in",
					Value: allowedApps,
				},
			},
		},
	}
	ok, err := repo.Exists(ctx, config.TokenDB, config.TokenColl, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotAuthorized
		}
		return err
	}
	if !ok {
		return ErrNotAuthorized
	}
	return nil
}
