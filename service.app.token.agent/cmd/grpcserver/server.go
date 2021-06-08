package grpcserver

import (
	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/fetching"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/initializing"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/modifying"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/validating"
)

type AppTokenServer struct {
	proto.UnimplementedAppTokenServer
	initService     initializing.Service
	modifySevice    modifying.Service
	fetchService    fetching.Service
	validateService validating.Service
}

// NewAppTokenServer creates all storage and domain dependencies binding them to the grpc-server
// and registers the grpc server implementing the grpc endpoints
func NewAppTokenServer(
	initService initializing.Service,
	modifyService modifying.Service,
	fetchService fetching.Service,
	validateService validating.Service) (*AppTokenServer, error) {

	return &AppTokenServer{
		initService:     initService,
		modifySevice:    modifyService,
		fetchService:    fetchService,
		validateService: validateService,
	}, nil
}

func (server AppTokenServer) mustEmbedUnimplementedAppTokenServer() {}
