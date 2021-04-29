package app

import (
	"context"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	tokenSrv "github.com/KonstantinGasser/datalab/protobuf/token_service"
	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/service_app/pkg/config"
	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
)

const (
	// DB, Collection names
	appDatabase   = "datalab_app"
	appCollection = "app"
)

// App describes what you can do with the App service
type App interface {
	// Create inserts a new app instance into a given storage after checking the app follows all
	// guide lines
	Create(ctx context.Context, stroage storage.Storage, appItem AppItem) errors.ErrApi

	// GetApps returns all apps created by a given user (forUUID)
	GetList(ctx context.Context, stroage storage.Storage, forUUID string) ([]AppItemLight, errors.ErrApi)

	// GetByID returns a specific app by its app uuid
	Get(ctx context.Context, stroage storage.Storage, userSrvice userSrv.UserClient, appUUID, callerUUID string) (*appSrv.ComplexApp, errors.ErrApi)

	// DeleteApp hard deletes an app from the mongo db - no setbacks
	Delete(ctx context.Context, stroage storage.Storage, appUUID, callerUUID, orgnAndApp string) errors.ErrApi

	// AddMember appends the member list of an app with the given member
	AddMember(ctx context.Context, storage storage.Storage, ownerUUID, appUUID string, member []string) errors.ErrApi

	// UpdateConfig updates the app configs for funnel, campaign and btn_time
	UpdateConfig(ctx context.Context, storage storage.Storage, cfg config.Cfgs, updateFlag string) errors.ErrApi

	// GenerateToken prepares all data required in order to generate an app token for the client library
	GetTokenClaims(ctx context.Context, storage storage.Storage, tokenSrv tokenSrv.TokenClient, appUUID, callerUUID, orgnAndApp string) (string, errors.ErrApi)

	// HasPermissions verifies that the request caller is allowed to work with the app resource
	HasPermissions(ctx context.Context, storage storage.Storage, callerUUID, appUUID string) errors.ErrApi
}

// app implements the App interface
type app struct{}

// NewApp returns a new app implementing the App interface
func NewApp() App {
	return &app{}
}

// AppItem represents one App in the database
type AppItem struct {
	// mongoDB pk (document key)
	UUID        string   `bson:"_id" required:"yes"`
	AppName     string   `bson:"name" required:"yes"`
	URL         string   `bson:"url" required:"yes"`
	OwnerUUID   string   `bson:"owner_uuid" required:"yes"`
	OrgnDomain  string   `bson:"orgn_domain" required:"yes"`
	Description string   `bson:"description"`
	Member      []string `bson:"member"`
	// Settings       []string    `bson:"setting" required:"yes" min:"1"`
	AppToken       string      `bson:"app_token"`
	Configurations config.Cfgs `bson:"app_config"`
	// OrgnAndAppHash is required to verify the generation of an app token
	// and the deletion of an app
	OrgnAndAppHash string `bson:"orgn_and_app_hash"`
}

// AppItemLight is a minimum representation of an application
type AppItemLight struct {
	// mongoDB pk (document key)
	UUID    string `bson:"_id" required:"yes"`
	AppName string `bson:"name" required:"yes"`
}

// // matchAppHash verifies that the request with domain name and app name matches with the database records
// // and that the request caller is the owner of the app
// func (app app) matchAppHash(ctx context.Context, storage storage.Storage, appUUID, callerUUID, domainAndName string) (bool, error) {
// 	query := bson.M{"_id": appUUID, "owner_uuid": callerUUID}

// 	var appData bson.M
// 	if err := storage.FindOne(ctx, appDatabase, appCollection, query, &appData); err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return false, errors.New("could not find any related documents for the given arguments")
// 		}
// 		return false, err
// 	}

// 	if _, ok := appData["orgn_and_app_hash"].(string); !ok {
// 		return false, errors.New("could not verify request to create app token")
// 	}

// 	requestHash := hash.Sha256([]byte(domainAndName)).String()
// 	if appData["orgn_and_app_hash"] != requestHash {
// 		return false, nil
// 	}
// 	return true, nil
// }
