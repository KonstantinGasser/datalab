package get

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/config"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFound = fmt.Errorf("could not find any token data")
)

func Token(ctx context.Context, repo repo.Repo, appUuid string) (*common.AppTokenInfo, error) {
	var token types.AppToken
	err := repo.FindOne(ctx, config.TokenDB, config.TokenColl, bson.M{"_id": appUuid}, &token)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &common.AppTokenInfo{Token: token.AppToken, Exp: token.Exp.Unix()}, nil
}
