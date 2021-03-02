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

func (srv UserService) mustEmbedUnimplementedUserServiceServer() {}
