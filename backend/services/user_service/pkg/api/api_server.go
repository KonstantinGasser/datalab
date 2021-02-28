package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/repository"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/user"
	"github.com/gofrs/uuid"
)

const (
	queryTimeout = time.Second * 5
)

type UserService struct {
	userSrv.UnimplementedUserServiceServer
	// *** Service Dependencies ***
	mongoClient *repository.MongoClient
	user        user.User
}

// NewUserService returns a pointer to a new UserService with all its
// dependencies
func NewUserService() (*UserService, error) {
	mongoC, err := repository.NewMongoClient("mongodb://userDB:secure@192.168.0.179:27017")
	log.Printf("could not create mongoDB client: %v\n", err)
	if err != nil {
		return nil, fmt.Errorf("could not create mongoDB client: %v", err)
	}
	return &UserService{
		mongoClient: mongoC,
		user:        user.User{},
	}, nil
}

// CreateUser receives the grpc request and handles user registration
func (srv UserService) CreateUser(ctx context.Context, request *userSrv.CreateUserRequest) (*userSrv.CreateUserResponse, error) {
	// generate random (NewV4()) user id for user and as _id
	// for collection in mongoDB
	UUID, err := uuid.NewV4()
	if err != nil {
		log.Printf("[userService.CreateUser] could not generate UUID for user: %v", err)
		return &userSrv.CreateUserResponse{
			StatusCode: http.StatusInternalServerError,
			Msg:        fmt.Sprintf("could not generate UUID for user: %v", err),
		}, fmt.Errorf("could not generate UUID for user: %v", err)
	}
	// insert new user in database
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	status, err := srv.user.Insert(ctx, srv.mongoClient, user.DBUser{
		UUID:       UUID.String(),
		Username:   request.GetUsername(),
		Password:   request.GetPassword(),
		OrgnDomain: request.GetOrgnDomain(),
	})
	if err != nil {
		return &userSrv.CreateUserResponse{
			StatusCode: int32(status),
			Msg:        err.Error(),
		}, nil
	}
	return &userSrv.CreateUserResponse{
		StatusCode: int32(status),
		Msg:        "user added to system",
	}, nil
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
