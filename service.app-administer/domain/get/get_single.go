package get

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFound = fmt.Errorf("could not find any related app")
)

func Single(ctx context.Context, repo repo.Repo, in *proto.GetRequest) (*common.AppInfo, error) {

	var app types.AppInfo
	err := repo.FindOne(ctx, config.AppDB, config.AppColl, bson.M{"_id": in.GetAppUuid()}, &app)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &common.AppInfo{
		Uuid:        app.Uuid,
		Name:        app.AppName,
		URL:         app.URL,
		Description: app.Description,
		Member:      app.Member,
	}, nil
}
