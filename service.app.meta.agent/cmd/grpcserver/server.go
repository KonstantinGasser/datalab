package grpcserver

import (
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/creating"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/fetching"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/inviting"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/updating"
)

type AppMetaServer struct {
	proto.UnimplementedAppMetaServer
	createSerivce creating.Service
	fechtService  fetching.Service
	inviteService inviting.Service
	updateService updating.Service
}

func NewAppMetaServer(
	createSerivce creating.Service,
	fechtService fetching.Service,
	inviteService inviting.Service,
	updateService updating.Service,
) *AppMetaServer {
	return &AppMetaServer{
		createSerivce: createSerivce,
		fechtService:  fechtService,
		inviteService: inviteService,
		updateService: updateService,
	}
}

func (server AppMetaServer) mustEmbedUnimplementedAppMetaServer() {}
