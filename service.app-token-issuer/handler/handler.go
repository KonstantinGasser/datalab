package handler

import (
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
)

type Handler struct {
	proto.UnimplementedAppTokenIssuerServer
	domain domain.AppTokenIssuer
	// *** Service Dependencies ***
}

func NewHandler(domain domain.AppTokenIssuer) *Handler {
	return &Handler{
		domain: domain,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (handler Handler) mustEmbedUnimplementedAppTokenIssuerServer() {}
