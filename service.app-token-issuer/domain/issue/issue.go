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
	"go.mongodb.org/mongo-driver/mongo"
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
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	var newExp = time.Now().Add(tokenExp)
	// logic for when app token already has been created
	if stored.AppToken != "" {
		if ok := override(stored.Exp); !ok {
			return nil, ErrTokenStillValid
		}
		stored.Exp = newExp
		token, err := jwts.Generate(in.GetAppUuid(), in.GetAppOrigin(), in.GetAppHash(), stored.Exp)
		if err != nil {
			return nil, err
		}
		stored.AppToken = token
		_, err = repo.UpdateOne(ctx, config.TokenDB, config.TokenColl, bson.M{"_id": stored.AppUuid}, bson.D{
			{
				Key:   "$set",
				Value: stored,
			},
		}, false)
		if err != nil {
			return nil, err
		}
		return &common.AppTokenInfo{
			Token: stored.AppToken,
			Exp:   stored.Exp.Unix(),
		}, nil
	}

	// logic for when no app token is present in database
	var newToken = types.AppToken{
		AppUuid:   in.GetAppUuid(),
		AppHash:   in.GetAppHash(),
		AppOrigin: in.GetAppOrigin(),
		Exp:       newExp,
	}
	token, err := jwts.Generate(in.GetAppUuid(), in.GetAppOrigin(), in.GetAppHash(), newExp)
	if err != nil {
		return nil, err
	}
	newToken.AppToken = token
	err = repo.InsertOne(ctx, config.TokenDB, config.TokenColl, newToken)
	if err != nil {
		return nil, err
	}
	return &common.AppTokenInfo{
		Token: newToken.AppToken,
		Exp:   newToken.Exp.Unix(),
	}, nil
}
