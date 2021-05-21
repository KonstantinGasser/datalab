package issue

import (
	"context"
	"fmt"
	"time"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/config"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/jwts"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	tokenExp = time.Hour * 24 * 7
)

var (
	ErrTokenStillValid = fmt.Errorf("app token is still valid. Thus can not be re-generated")
)

func Token(ctx context.Context, repo repo.Repo, in *proto.IssueRequest) (*common.AppTokenInfo, error) {

	var stored *types.AppToken = &types.AppToken{}
	err := repo.FindOne(ctx, config.TokenDB, config.TokenColl, bson.M{"_id": in.GetAppUuid()}, stored)
	if err != nil {
		return nil, err
	}

	// logic for when app token already has been created
	// new app token can only be issued if current one has expired
	if stored.AppToken != "" {
		if ok := override(stored.Exp); !ok {
			return nil, ErrTokenStillValid
		}
	}

	var newExp = time.Now().Add(tokenExp)
	token, err := jwts.Generate(in.GetAppUuid(), in.GetAppOrigin(), in.GetAppHash(), newExp)
	if err != nil {
		return nil, err
	}
	_, err = repo.UpdateOne(ctx, config.TokenDB, config.TokenColl, bson.M{"_id": stored.AppUuid}, bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"app_token":  token,
				"token_exp":  newExp,
				"app_origin": in.GetAppOrigin(),
			},
		},
	}, false)
	if err != nil {
		return nil, err
	}
	return &common.AppTokenInfo{
		Token: token,
		Exp:   newExp.Unix(),
	}, nil
}
