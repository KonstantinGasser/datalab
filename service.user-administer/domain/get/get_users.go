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
	ErrNoUsersFound = fmt.Errorf("could not find any users from uuid list")
)

func Users(ctx context.Context, repo repo.Repo, in *proto.GetListRequest) ([]*common.UserInfo, error) {

	var foundUsers []types.UserInfo
	err := repo.FindMany(ctx, config.UserDB, config.UserColl, bson.D{
		{
			Key: "_id",
			Value: bson.M{
				"$in": in.GetUuidList(),
			},
		},
	}, &foundUsers)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNoUsersFound
		}
		return nil, err
	}
	fmt.Println("uuids: ", foundUsers)

	var users = make([]*common.UserInfo, len(foundUsers))
	for i, item := range foundUsers {
		users[i] = &common.UserInfo{
			Uuid:          item.Uuid,
			Username:      item.Username,
			FirstName:     item.FirstName,
			LastName:      item.LastName,
			OrgnDomain:    item.OrgnDomain,
			OrgnPosition:  item.OrgnPosition,
			ProfileImgUrl: item.ProfileImgURL,
		}
	}
	return users, nil
}
