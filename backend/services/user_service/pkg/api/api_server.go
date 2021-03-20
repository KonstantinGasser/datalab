package api

import (
	"time"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/storage"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/user"
)

const (
	queryTimeout = time.Second * 5
)

// UserService implements all the methods required by the grpc.UserServiceServer
// and embeds all the required dependencies
type UserService struct {
	userSrv.UnimplementedUserServiceServer
	// *** Service Dependencies ***
	storage storage.Storage
	user    user.User
}

// NewUserService returns a pointer to a new UserService with all its
// dependencies
func NewUserService(storage storage.Storage) *UserService {
	return &UserService{
		// *** Storage Dependencies ***
		storage: storage,
		// *** Service Dependencies
		user: user.NewUser(),
	}
}

// mustEmbedUnimplementedUserServiceServer as by the current grpc version the
// server must implement this method in order for it to act as a grpc.UserServiceServer
func (srv UserService) mustEmbedUnimplementedUserServiceServer() {}
