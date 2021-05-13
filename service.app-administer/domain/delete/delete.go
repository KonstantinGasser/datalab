package delete

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/hasher"
	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/permissions"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrNoPermissions = fmt.Errorf("App-Hash does not match with records")
)

// App deletes an existing App. It returns the App in case a compensating action needs
// to be performed
func App(ctx context.Context, repo repo.Repo, in *proto.DeleteRequest) (*types.AppInfo, error) {
	hash := hasher.Build(in.GetAppName(), in.GetOrgnName())
	if err := permissions.IsCorrectHash(ctx, repo, in.GetAppUuid(), hash); err != nil {
		if err == permissions.ErrNotAuthorized {
			return nil, ErrNoPermissions
		}
		return nil, err
	}

	var app types.AppInfo
	err := repo.FindOne(ctx, config.AppDB, config.AppColl, bson.M{"_id": in.GetAppUuid()}, &app)
	if err != nil {
		return nil, err
	}

	err = repo.DeleteOne(ctx, config.AppDB, config.AppColl, bson.M{"_id": in.GetAppUuid()})
	if err != nil {
		return nil, err
	}
	return &app, nil
}
