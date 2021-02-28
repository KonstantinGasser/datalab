package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/repository"
	"github.com/gofrs/uuid"
)

type UserService struct {
	userSrv.UnimplementedUserServiceServer
	// *** Service Dependencies ***
	mongoClient *repository.MongoClient
}

func NewUserService() (*UserService, error) {
	mongoC, err := repository.NewMongoClient("mongodb://userDB:secure@192.168.0.179:27017")
	if err != nil {
		return nil, fmt.Errorf("could not create mongoDB client: %v", err)
	}

	return &UserService{
		mongoClient: mongoC,
	}, nil
}

func (srv UserService) CreateUser(ctx context.Context, request *userSrv.CreateUserRequest) (*userSrv.CreateUserResponse, error) {
	log.Printf("received: %v\n", request)
	// generate random (NewV4()) user id for user and as _id
	// for collection in mongoDB
	UUID, err := uuid.NewV4()
	if err != nil {
		return &userSrv.CreateUserResponse{
			StatusCode: http.StatusInternalServerError,
			Msg: fmt.Sprintf("could not generate UUID for user: %v", err)
		},
	}
	// insert new user in database
	srv.mongoClient.InsertUser(repository.MongoUser{
		UUID: UUID,
		Username: request.Username,
		// needs some salt and pepper...
		Password: request.Password, 
		OrgnDomain: request.OrgnDomain,
	})

}
func (srv UserService) AuthUser(ctx context.Context, request *userSrv.AuthRequest) (*userSrv.AuthResponse, error) {
	log.Printf("GRPC Request: %v", request)
	return &userSrv.AuthResponse{
		StatusCode:    http.StatusOK,
		Msg:           "first grpc stuff I am doing",
		Authenticated: true,
		User: &userSrv.AuthenticatedUser{
			Username: "KonstantinGasser",
		},
	}, nil
}

func (srv UserService) mustEmbedUnimplementedUserServiceServer() {}
