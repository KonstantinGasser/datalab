package handler

import (
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain"
	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
)

type Handler struct {
	proto.UnimplementedUserAuthenticationServer
	domain domain.UserAuthLogic
	// *** Service Dependencies ***
}

// NewHandler returns a new *Handler wrapping the domain logic
func NewHandler(domain domain.UserAuthLogic) *Handler {
	return &Handler{
		domain: domain,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (handler Handler) mustEmbedUnimplementedUserAuthenticationServer() {}
