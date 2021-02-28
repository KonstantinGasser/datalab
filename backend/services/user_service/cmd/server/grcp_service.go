package server

// import (
// 	"context"
// 	"log"
// 	"net/http"

// 	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
// )

// type userService struct {
// 	userSrv.UnimplementedUserServiceServer
// }

// func (srv userService) CreateUser(ctx context.Context, request *userSrv.CreateUserRequest) (*userSrv.CreateUserResponse, error) {
// 	log.Printf("recived: %v\n", request)
// 	return &userSrv.CreateUserResponse{
// 		StatusCode: http.StatusOK,
// 		Msg:        "you're all good, mate!",
// 	}, nil
// }
// func (srv userService) AuthUser(ctx context.Context, request *userSrv.AuthRequest) (*userSrv.AuthResponse, error) {
// 	log.Printf("GRPC Request: %v", request)
// 	return &userSrv.AuthResponse{
// 		StatusCode:    http.StatusOK,
// 		Msg:           "first grpc stuff I am doing",
// 		Authenticated: true,
// 		User: &userSrv.AuthenticatedUser{
// 			Username: "KonstantinGasser",
// 		},
// 	}, nil
// }

// func (srv userService) mustEmbedUnimplementedUserServiceServer() {}
