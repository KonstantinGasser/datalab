package permissions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotAuthorized  = fmt.Errorf("caller is not authorized to perform the action")
	ErrNotFound       = fmt.Errorf("could not find app information")
	ErrNoAppHashFound = fmt.Errorf("could not find any app hash for app uuid")
	ErrNoAppsFound    = fmt.Errorf("could not find any apps the user is allowed to have access to")
)

type Permission struct {
	permissionRepo Repository
}

func New(repo Repository) Permission {
	return Permission{
		permissionRepo: repo,
	}
}

// IsOwner checks if a given user is the owner (has read/write access) of the data
func (p Permission) IsOwner(ctx context.Context, repo Repository, callerUuid, appUuid string) errors.ErrApi {

	err := p.permissionRepo.HasRWAccess(ctx, appUuid, callerUuid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New(http.StatusUnauthorized, ErrNotAuthorized, "Must be Owner of App for this action")
		}
		return errors.New(http.StatusUnauthorized, err, "Could not check permissions")
	}

	return nil
}

// func (p Permission) CanAccess(ctx context.Context, repo repo.Repo, permissions *common.UserTokenClaims, appUuid string) error {
// 	fmt.Printf("permissions: %+v\n", permissions)
// 	if permissions.GetPermissions().GetApps() == nil || len(permissions.GetPermissions().GetApps()) == 0 {
// 		return ErrNotAuthorized
// 	}
// 	var allowedApps []string
// 	for _, item := range permissions.GetPermissions().GetApps() {
// 		allowedApps = append(allowedApps, item.GetAppUuid())
// 	}
// 	filter := bson.D{
// 		{
// 			Key: "_id",
// 			Value: bson.D{
// 				{
// 					Key:   "$in",
// 					Value: allowedApps,
// 				},
// 			},
// 		},
// 	}
// 	ok, err := repo.Exists(ctx, config.TokenDB, config.TokenColl, filter)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return ErrNotAuthorized
// 		}
// 		return err
// 	}
// 	if !ok {
// 		return ErrNotAuthorized
// 	}
// 	return nil
// }
