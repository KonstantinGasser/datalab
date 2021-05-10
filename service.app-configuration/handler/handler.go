package handler

import (
	"github.com/KonstantinGasser/datalab/service.app-configuration/domain"
	"github.com/KonstantinGasser/datalab/service.app-configuration/proto"
)

type Handler struct {
	proto.UnimplementedAppConfigurationServer
	domain domain.AppConfig
	// *** Service Dependencies ***
}

func NewHandler(domain domain.AppConfig) *Handler {
	return &Handler{
		domain: domain,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (handler Handler) mustEmbedUnimplementedAppConfigurationServer() {}
