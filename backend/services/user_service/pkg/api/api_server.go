package api

import (
	"time"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/repository"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/user"
)

const (
	queryTimeout = time.Second * 5
)

// UserService implements all the methods required by the grpc.UserServiceServer
// and embeds all the required dependencies
type UserService struct {
	userSrv.UnimplementedUserServiceServer
	user user.User
	// *** Service Dependencies ***
	mongoClient *repository.MongoClient
}

// NewUserService returns a pointer to a new UserService with all its
// dependencies
func NewUserService(mongoC *repository.MongoClient) *UserService {
	return &UserService{
		mongoClient: mongoC,
		user:        user.User{},
	}
}

// mustEmbedUnimplementedUserServiceServer as by the current grpc version the
// server must implement this method in order for it to act as a grpc.UserServiceServer
func (srv UserService) mustEmbedUnimplementedUserServiceServer() {}
