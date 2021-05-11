package initialize

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app-configuration/config"
	"github.com/KonstantinGasser/datalab/service.app-configuration/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/service.app-configuration/repo"
)

func Configs(ctx context.Context, repo repo.Repo, in *proto.InitRequest) error {
	var cfg = types.ConfigInfo{AppUuid: in.GetForApp()}
	if err := repo.InsertOne(ctx, config.CfgDB, config.CfgColl, cfg); err != nil {
		return err
	}
	return nil
}
