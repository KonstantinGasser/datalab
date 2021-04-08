package api

import (
	tokenSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/token_service"
)

type TokenServer struct {
	tokenSrv.UnimplementedTokenServer
	// ** Service Dependencies ***
}

// NewTokenService creates and returns a new TokenService
func NewTokenServer() TokenServer {
	srv := TokenServer{}
	return srv
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv TokenServer) mustEmbedUnimplementedTokenServiceServer() {}
