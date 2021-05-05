package api

import (
	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	apptokenSrv "github.com/KonstantinGasser/datalab/protobuf/apptoken_service"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/storage"
)

type AppTokenServer struct {
	apptokenSrv.UnimplementedAppTokenServer
	// ** Service Dependencies ***
	app     appSrv.AppClient
	storage storage.Storage
}

// NewTokenService creates and returns a new TokenService
func NewAppTokenServer(app appSrv.AppClient, storage storage.Storage) *AppTokenServer {
	return &AppTokenServer{
		app:     app,
		storage: storage,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv AppTokenServer) mustEmbedUnimplementedAppTokenServer() {}
