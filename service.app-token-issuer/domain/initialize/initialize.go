package initialize

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/config"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/repo"
)

func AppToken(ctx context.Context, repo repo.Repo, appUuid, appHash, ownerUuid string) error {

	var appToken = types.AppToken{
		AppUuid:  appUuid,
		AppHash:  appHash,
		AppOwner: ownerUuid,
	}
	err := repo.InsertOne(ctx, config.TokenDB, config.TokenColl, appToken)
	if err != nil {
		return err
	}
	return nil
}
