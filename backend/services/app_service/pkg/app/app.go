package app

import (
	"context"

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
	CreateApp(ctx context.Context, stroage storage.Storage, appItem AppItem) (int, error)
	// GetApps returns all apps created by a given user (forUUID)
	GetApps(ctx context.Context, stroage storage.Storage, forUUID string) (int, []AppItemLight, error)
	// GetByID returns a specific app by its app uuid
	GetByID(ctx context.Context, stroage storage.Storage, appUUID string) (int, AppItem, error)
	// DeleteApp hard deletes an app from the mongo db - no setbacks
	DeleteApp(ctx context.Context, stroage storage.Storage, appUUID string) (int, error)
	// AddMember appends the member list of an app with the given member
	AddMember(ctx context.Context, storage storage.Storage, ownerUUID, appUUID string, member []string) (int, error)
}

// NewApp returns a new app implementing the App interface
func NewApp() App {
	return &app{}
}
