package initialize

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app-configuration/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/service.app-configuration/repo"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/config"
)

func Configs(ctx context.Context, repo repo.Repo, in *proto.InitRequest) error {
	var cfg = types.ConfigInfo{AppUuid: in.GetForApp()}
	if err := repo.InsertOne(ctx, config.TokenDB, config.TokenColl, cfg); err != nil {
		return err
	}
	return nil
}
