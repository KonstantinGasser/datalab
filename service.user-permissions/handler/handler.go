package handler

import (
	"github.com/KonstantinGasser/datalab/service.user-permissions/domain"
	"github.com/KonstantinGasser/datalab/service.user-permissions/proto"
)

type Handler struct {
	proto.UnimplementedUserPermissionsServer
	// *** Service Dependencies ***
	domain domain.UserPermissions
}

func NewHandler(domain domain.UserPermissions) *Handler {
	return &Handler{
		domain: domain,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (handler Handler) mustEmbedUnimplementedUserPermissionsServer() {}
