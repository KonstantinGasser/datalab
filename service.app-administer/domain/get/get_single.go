package get

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFound = fmt.Errorf("could not find any related app")
)

// Single looks up the app data for one given app uuid
func Single(ctx context.Context, repo repo.Repo, appUuid string) (*common.AppInfo, error) {

	var app types.AppInfo
	err := repo.FindOne(ctx, config.AppDB, config.AppColl, bson.M{"_id": appUuid}, &app)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var member = make([]*common.AppMember, len(app.Invites))
	for i, item := range app.Invites {
		member[i] = &common.AppMember{Uuid: item.Uuid, Status: int32(item.Status)}
	}
	return &common.AppInfo{
		Uuid:        app.Uuid,
		Name:        app.AppName,
		URL:         app.URL,
		Description: app.Description,
		Member:      member,
		Owner:       app.OwnerUuid,
	}, nil
}
