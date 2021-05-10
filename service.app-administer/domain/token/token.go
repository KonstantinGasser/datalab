package token

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service.app-administer/domain/hasher"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/permissions"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
)

var (
	ErrNotAuthorized = fmt.Errorf("caller is not authorized to perform the action")
)

func MayAcquire(ctx context.Context, repo repo.Repo, in *proto.MayAcquireTokenRequest) (bool, error) {

	if err := permissions.IsOwner(ctx, repo, in.GetCallerUuid(), in.GetAppUuid()); err != nil {
		if err == permissions.ErrNotAuthorized {
			return false, ErrNotAuthorized
		}
		return false, err
	}
	if err := hasher.Compare(ctx, repo, in.GetAppHash(), in.GetAppUuid()); err != nil {
		if err == hasher.ErrHashMisMatch {
			return false, ErrNotAuthorized
		}
		return false, err
	}
	return true, nil
}
