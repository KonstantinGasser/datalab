package grpcserver

import (
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/adding"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/fetching"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/users/authenticating"
)

type UserAuthServer struct {
	proto.UnimplementedUserAuthenticationServer
	authService  authenticating.Service
	addService   adding.Service
	fetchService fetching.Service
}

func NewUserAuthServer(
	authService authenticating.Service,
	addService adding.Service,
	fetchService fetching.Service,
) *UserAuthServer {
	return &UserAuthServer{
		authService:  authService,
		addService:   addService,
		fetchService: fetchService,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (server UserAuthServer) mustEmbedUnimplementedUserAuthenticationServer() {}
