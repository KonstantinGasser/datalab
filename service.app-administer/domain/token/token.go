package token

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service.app-administer/domain/permissions"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
)

var (
	ErrNotAuthorized = fmt.Errorf("caller is not authorized to perform the action")
	ErrNotFound      = fmt.Errorf("could not find any related app")
)

func MayAcquire(ctx context.Context, repo repo.Repo, in *proto.MayAcquireTokenRequest) (bool, error) {
	if err := permissions.IsOwner(ctx, repo, in.GetUserClaims().GetUuid(), in.GetAppUuid()); err != nil {
		if err == permissions.ErrNotAuthorized {
			return false, ErrNotAuthorized
		}
		return false, err
	}
	if err := permissions.IsCorrectHash(ctx, repo, in.GetAppUuid(), in.GetAppHash()); err != nil {
		if err == permissions.ErrNotFound {
			return false, ErrNotFound
		}
		if err == permissions.ErrNotAuthorized {
			return false, ErrNotAuthorized
		}
		return false, err
	}
	return true, nil
}
