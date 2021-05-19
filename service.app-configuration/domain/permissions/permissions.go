package permissions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-configuration/config"
	"github.com/KonstantinGasser/datalab/service.app-configuration/errors"
	"github.com/KonstantinGasser/datalab/service.app-configuration/repo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrNotAuthorized  = fmt.Errorf("caller is not authorized to perform the action")
	ErrNotFound       = fmt.Errorf("could not find app information")
	ErrNoAppHashFound = fmt.Errorf("could not find any app hash for app uuid")
	ErrNoAppsFound    = fmt.Errorf("could not find any apps the user is allowed to have access to")
)

func CanAccess(ctx context.Context, repo repo.Repo, permissions *common.UserTokenClaims, appUuid string) errors.ErrApi {

	var allowedApps []string
	for _, item := range permissions.GetPermissions().GetApps() {
		allowedApps = append(allowedApps, item.GetAppUuid())
	}
	fmt.Println(permissions)
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
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not check if owner of app",
			Err:    err,
		}
	}
	if !ok {
		return errors.ErrAPI{
			Status: http.StatusUnauthorized,
			Msg:    "User has no permissions",
			Err:    err,
		}
	}
	return nil
}
