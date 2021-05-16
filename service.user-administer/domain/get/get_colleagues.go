package get

import (
	"context"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.user-administer/config"
	"github.com/KonstantinGasser/datalab/service.user-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Colleaues(ctx context.Context, repo repo.Repo, userUuid string) ([]*common.UserInfo, error) {

	baseFilter := bson.M{"_id": userUuid}
	var baseUser types.UserInfo
	err := repo.FindOne(ctx, config.UserDB, config.UserColl, baseFilter, &baseUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNoUserFound
		}
		return nil, err
	}

	// queries for all uses with the organization domain as the provided userUuid has
	// ignores passed in userUuid
	colleaguesFilter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.M{"orgn_domain": baseUser.OrgnDomain},
				bson.D{
					{
						Key:   "_id",
						Value: bson.M{"$ne": baseUser.Uuid},
					},
				},
			},
		},
	}
	var foundColleauges []types.UserInfo
	err = repo.FindMany(ctx, config.UserDB, config.UserColl, colleaguesFilter, &foundColleauges)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNoUsersFound
		}
		return nil, err
	}
	var colleagues = make([]*common.UserInfo, len(foundColleauges))
	for i, item := range foundColleauges {
		colleagues[i] = &common.UserInfo{
			Uuid:          item.Uuid,
			Username:      item.Username,
			FirstName:     item.FirstName,
			LastName:      item.LastName,
			OrgnPosition:  item.OrgnPosition,
			OrgnDomain:    item.OrgnDomain,
			ProfileImgUrl: item.ProfileImgURL,
		}
	}
	return colleagues, nil
}
