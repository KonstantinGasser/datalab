package get

import (
	"context"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/config"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

func Token(ctx context.Context, repo repo.Repo, appUuid string) (*common.AppTokenInfo, error) {
	var token types.AppToken
	err := repo.FindOne(ctx, config.TokenDB, config.TokenColl, bson.M{"_id": appUuid}, &token)
	if err != nil {
		return nil, err
	}
	return &common.AppTokenInfo{Token: token.AppToken, Exp: token.Exp.Unix()}, nil
}
