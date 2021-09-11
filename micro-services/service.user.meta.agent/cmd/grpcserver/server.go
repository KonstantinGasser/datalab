package grpcserver

import (
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/creating"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/fetching"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/updating"
)

type UserMetaServer struct {
	proto.UnimplementedUserMetaServer
	createService creating.Service
	updateService updating.Service
	fetchService  fetching.Service
}

func NewUserMetaServer(
	createService creating.Service,
	updateService updating.Service,
	fetchService fetching.Service,
) *UserMetaServer {
	return &UserMetaServer{
		createService: createService,
		updateService: updateService,
		fetchService:  fetchService,
	}
}

func (server UserMetaServer) mustEmbedUnimplementedUserMetaServer() {}
