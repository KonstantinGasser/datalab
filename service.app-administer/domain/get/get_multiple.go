package get

import (
	"context"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

func Multiple(ctx context.Context, repo repo.Repo, in *proto.GetListRequest) ([]*common.AppMetaInfo, error) {

	var data []types.AppMetaInfo
	err := repo.FindMany(ctx, config.AppDB, config.AppColl, bson.M{"owner_uuid": in.GetCallerUuid()}, &data)
	if err != nil {
		return nil, ErrNotFound
	}

	var apps = make([]*common.AppMetaInfo, len(data))
	for i, item := range data {
		apps[i] = &common.AppMetaInfo{Uuid: item.Uuid, Name: item.AppName}
	}
	return apps, nil
}
