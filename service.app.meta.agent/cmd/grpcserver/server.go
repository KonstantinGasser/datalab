package grpcserver

import (
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/creating"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/fetching"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/inviting"
)

type AppMetaServer struct {
	proto.UnimplementedAppMetaServer
	createSerivce creating.Service
	fechtService  fetching.Service
	inviteService inviting.Service
}

func NewAppMetaServer(
	createSerivce creating.Service,
	fechtService fetching.Service,
	inviteService inviting.Service,
) *AppMetaServer {
	return &AppMetaServer{
		createSerivce: createSerivce,
		fechtService:  fechtService,
		inviteService: inviteService,
	}
}

func (server AppMetaServer) mustEmbedUnimplementedAppMetaServer() {}
