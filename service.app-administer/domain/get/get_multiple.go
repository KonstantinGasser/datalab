package get

import (
	"context"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

// Multiple looks up all apps a user can access given its permissions
func Multiple(ctx context.Context, repo repo.Repo, appUuids ...string) ([]*common.AppMetaInfo, error) {

	filter := bson.D{
		{
			Key: "_id",
			Value: bson.D{
				{
					Key:   "$in",
					Value: appUuids,
				},
			},
		},
	}
	var data []types.AppMetaInfo
	err := repo.FindMany(ctx, config.AppDB, config.AppColl, filter, &data)
	if err != nil {
		return nil, ErrNotFound
	}

	var apps = make([]*common.AppMetaInfo, len(data))
	for i, item := range data {
		apps[i] = &common.AppMetaInfo{Uuid: item.Uuid, Name: item.AppName}
	}
	return apps, nil
}
