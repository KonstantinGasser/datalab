package api

import (
	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	"github.com/KonstantinGasser/datalab/service_config/pkg/config"
	"github.com/KonstantinGasser/datalab/service_config/pkg/storage"
)

type ConfigServer struct {
	configSrv.UnimplementedConfigServer
	// ** Service Dependencies ***
	config  config.Config
	storage storage.Storage
}

// NewTokenService creates and returns a new TokenService
func NewConfigServer(cfg config.Config, storage storage.Storage) *ConfigServer {
	return &ConfigServer{
		config:  cfg,
		storage: storage,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv ConfigServer) mustEmbedUnimplementedConfigServer() {}
