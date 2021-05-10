package handler

import (
	"github.com/KonstantinGasser/datalab/service.user-administer/domain"
	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
)

type Handler struct {
	proto.UnimplementedUserAdministerServer
	domain domain.UserAdminLogic
	// *** Service Dependencies ***
}

func NewHandler(domain domain.UserAdminLogic) *Handler {
	return &Handler{
		domain: domain,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (handler Handler) mustEmbedUnimplementedUserAdministerServer() {}
