package permissions

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-configuration/config"
	"github.com/KonstantinGasser/datalab/service.app-configuration/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotAuthorized  = fmt.Errorf("caller is not authorized to perform the action")
	ErrNotFound       = fmt.Errorf("could not find app information")
	ErrNoAppHashFound = fmt.Errorf("could not find any app hash for app uuid")
	ErrNoAppsFound    = fmt.Errorf("could not find any apps the user is allowed to have access to")
)

func CanAccess(ctx context.Context, repo repo.Repo, permissions *common.UserTokenClaims, appUuid string) error {

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
	ok, err := repo.Exists(ctx, config.CfgDB, config.CfgColl, filter)
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
