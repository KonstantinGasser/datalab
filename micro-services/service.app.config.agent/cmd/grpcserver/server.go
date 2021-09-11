package grpcserver

import (
	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/fetching"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/initializing"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/modifying"
	libFetch "github.com/KonstantinGasser/datalab/service.app.config.agent/internal/libconfig/fetching"
)

type AppConfigServer struct {
	proto.UnimplementedAppConfigurationServer
	initService     initializing.Service
	modifyService   modifying.Service
	fetchService    fetching.Service
	libFetchService libFetch.Service
}

func NewAppConfigServer(
	initService initializing.Service,
	modifyService modifying.Service,
	fetchService fetching.Service,
	libFetchService libFetch.Service,
) *AppConfigServer {
	return &AppConfigServer{
		initService:     initService,
		modifyService:   modifyService,
		fetchService:    fetchService,
		libFetchService: libFetchService,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (server AppConfigServer) mustEmbedUnimplementedAppConfigurationServer() {}
