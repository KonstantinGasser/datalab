package create

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/hasher"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"github.com/KonstantinGasser/datalab/utils/unique"
)

// App creates a instance of types.App generates the App-Uuid and the App-Hash
// and inserts the new App into the database
func App(ctx context.Context, repo repo.Repo, in *proto.CreateRequest) (string, error) {

	appUuid, err := unique.UUID()
	if err != nil {
		return "", err
	}
	appHash := hasher.Build(in.GetName(), in.GetOrganization())
	app := types.AppInfo{
		Uuid:        appUuid,
		AppName:     in.GetName(),
		URL:         in.GetAppUrl(),
		OwnerUuid:   in.GetOwnerUuid(),
		OrgnDomain:  in.GetOrganization(),
		Description: in.GetDescription(),
		Member:      nil,
		AppHash:     appHash,
	}

	if err := repo.InsertOne(ctx, config.AppDB, config.AppColl, app); err != nil {
		return "", err
	}
	return appUuid, nil
}
