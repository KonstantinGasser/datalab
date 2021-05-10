package get

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.user-administer/config"
	"github.com/KonstantinGasser/datalab/service.user-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/service.user-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNoUserFound = fmt.Errorf("could not find any user")
)

func User(ctx context.Context, repo repo.Repo, in *proto.GetRequest) (*common.UserInfo, error) {

	var foundUser types.UserInfo
	err := repo.FindOne(ctx, config.UserDB, config.UserColl, bson.M{"_id": in.GetForUuid()}, &foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNoUserFound
		}
		return nil, err
	}

	var user = &common.UserInfo{
		Uuid:          foundUser.Uuid,
		Username:      foundUser.Username,
		FirstName:     foundUser.FirstName,
		LastName:      foundUser.LastName,
		OrgnDomain:    foundUser.OrgnDomain,
		OrgnPosition:  foundUser.OrgnPosition,
		ProfileImgUrl: foundUser.ProfileImgURL,
	}

	return user, nil
}
