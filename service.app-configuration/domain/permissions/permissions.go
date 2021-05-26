package permissions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
)

var (
	ErrNotAuthorized  = fmt.Errorf("caller is not authorized to perform the action")
	ErrNotFound       = fmt.Errorf("could not find app information")
	ErrNoAppHashFound = fmt.Errorf("could not find any app hash for app uuid")
	ErrNoAppsFound    = fmt.Errorf("could not find any apps the user is allowed to have access to")
)

type Permission struct {
	permissionRepo PermissionDao
}

func New(repo PermissionDao) Permission {
	return Permission{
		permissionRepo: repo,
	}
}

func (p Permission) HasRead(ctx context.Context, appUuid string, allowedReads []*common.AppPermission) errors.ErrApi {
	for _, item := range allowedReads {
		if item.AppUuid == appUuid {
			return nil
		}
	}
	return errors.New(http.StatusUnauthorized, ErrNotAuthorized, "User is not allowed to see App Config")
}
