package app

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/app_service"
	userSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
)

// why are they here?????? change this pls
const (
	// DB, Collection names
	appDatabase   = "datalabs_app"
	appCollection = "app"
)

// App describes what you can do with the App service
type App interface {
	// Create inserts a new app instance into a given storage after checking the app follows all
	// guide lines
	CreateApp(ctx context.Context, stroage storage.Storage, appItem AppItem) (int, error)

	// GetApps returns all apps created by a given user (forUUID)
	GetAppList(ctx context.Context, stroage storage.Storage, forUUID string) (int, []AppItemLight, error)

	// GetByID returns a specific app by its app uuid
	GetApp(ctx context.Context, stroage storage.Storage, userSrvice userSrv.UserClient, appUUID, callerUUID string) (int, *appSrv.ComplexApp, error)

	// DeleteApp hard deletes an app from the mongo db - no setbacks
	DeleteApp(ctx context.Context, stroage storage.Storage, appUUID string) (int, error)

	// AddMember appends the member list of an app with the given member
	AddMember(ctx context.Context, storage storage.Storage, ownerUUID, appUUID string, member []string) (int, error)

	// CanGenToken verifies, that the caller intentionally wants to create an app token by cross checking that
	// the correct domain/app-name was entered
	CanGenToken(ctx context.Context, storage storage.Storage, appUUID, callerUUID, domainAndName string) (int, bool, error)
}

// app implements the App interface
type app struct{}

// NewApp returns a new app implementing the App interface
func NewApp() App {
	return &app{}
}
