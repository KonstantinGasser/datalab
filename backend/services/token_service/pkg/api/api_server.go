package api

import (
	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
)

type TokenService struct {
	tokenSrv.UnimplementedTokenServiceServer
	// ** Service Dependencies ***
}

// NewTokenService creates and returns a new TokenService
func NewTokenService() TokenService {
	srv := TokenService{}
	return srv
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv TokenService) mustEmbedUnimplementedTokenServiceServer() {}
