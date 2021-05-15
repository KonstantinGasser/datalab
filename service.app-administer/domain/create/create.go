package create

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/hasher"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"github.com/KonstantinGasser/datalab/utils/unique"
	"github.com/KonstantinGasser/required"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrAppNameExists = fmt.Errorf("app name already exists but must be unique")
	ErrMissingFields = fmt.Errorf("app struct is missing mandatory fields")
)

// App creates a instance of types.App, generates the App-Uuid and the App-Hash
// and inserts the new App into the database
func App(ctx context.Context, repo repo.Repo, in *proto.CreateRequest) (string, error) {

	exists, err := repo.Exists(ctx, config.AppDB, config.AppColl,
		bson.M{"name": in.GetName(), "owner_uuid": in.GetOwnerUuid()})
	if err != nil && err != mongo.ErrNilDocument {
		return "", err
	}
	if exists {
		return "", ErrAppNameExists
	}
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
	// before inserting the app ensure that the assertion to check
	// if all mandatory fields does not fail - if so returns an ErrMissingFields
	if err := required.Atomic(&app); err != nil {
		return "", ErrMissingFields
	}
	logrus.Warn(app)
	if err := repo.InsertOne(ctx, config.AppDB, config.AppColl, app); err != nil {
		return "", err
	}
	return appUuid, nil
}
