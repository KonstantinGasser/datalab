package app

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
)

// why are they here?????? change this pls
const (
	// DB, Collection names
	dbName        = "datalabs_app"
	appCollection = "app"
)

// App describes what you can do with the App service
type App interface {
	CreateApp(ctx context.Context, mongo storage.Storage, req *appSrv.CreateAppRequest) (int, error)
	GetApps(ctx context.Context, mongo storage.Storage, req *appSrv.GetAppsRequest) ([]*appSrv.LightApp, error)
}
type app struct{}

// New returns a new app implementing the App interface
func New() App {
	return &app{}
}

// AppItem represents one App in the database ? do we need this? don't we have a def in the grpc already???
type AppItem struct {
	// mongoDB pk (document key)
	UUID       string `bson:"_id"`
	AppName    string `bson:"appName"`
	OwnerUUID  string `bson:"ownerUUID"`
	OrgnDomain string `bson:"orgnDomain"`
	// Member is a list of user_uuids mapped to this app
	Member   []string `bson:"member"`
	AppToken string   `bson:"appToken"`
}
