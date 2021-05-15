package initialize

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.user-permissions/config"
	"github.com/KonstantinGasser/datalab/service.user-permissions/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-permissions/proto"
	"github.com/KonstantinGasser/datalab/service.user-permissions/repo"
)

func Permissions(ctx context.Context, repo repo.Repo, in *proto.InitRequest) error {

	var permissions = types.Permissions{
		UserUuid: in.GetUserUuid(),
		UserOrgn: in.GetUserOrgn(),
		Apps:     nil,
	}
	err := repo.InsertOne(ctx, config.PermissonDB, config.PermissonColl, permissions)
	if err != nil {
		return err
	}
	return nil
}
