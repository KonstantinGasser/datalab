package handler

import (
	"github.com/KonstantinGasser/datalab/service.app-administer/domain"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	cfgsvc "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	aptissuer "github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	usersvc "github.com/KonstantinGasser/datalab/service.user-administer/proto"
)

// Handler implements the grpc Service interface
type Handler struct {
	proto.UnimplementedAppAdministerServer
	domain domain.AppAdmin
	// *** Service Dependencies ***
	userSvc   usersvc.UserAdministerClient
	configSvc cfgsvc.AppConfigurationClient
	tokenSvc  aptissuer.AppTokenIssuerClient
}

func NewHandler(domain domain.AppAdmin) *Handler {
	return &Handler{
		domain: domain,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (handler Handler) mustEmbedUnimplementedAppAdministerServer() {}
