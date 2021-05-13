package register

import (
	"context"
	"fmt"
	"strings"

	"github.com/KonstantinGasser/datalab/service.user-administer/repo"
	"github.com/KonstantinGasser/datalab/service.user-authentication/config"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/utils/hash"
	"github.com/KonstantinGasser/datalab/utils/unique"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrUserAlreadyExists = fmt.Errorf("user with this username already exists")
	ErrInvalidOrgnName   = fmt.Errorf("organization name must not contain a forward-slash")
)

// NewUser registers a new user in the database
func NewUser(ctx context.Context, repo repo.Repo, in *proto.RegisterRequest) (string, error) {
	if !orgnNameAllowed(in.GetOrganisation()) {
		return "", ErrInvalidOrgnName
	}

	exists, err := repo.Exists(ctx, config.UserAuthDB, config.UserAuthColl, bson.M{"username": in.GetUsername()})
	if err != nil {
		return "", err
	}
	if exists {
		return "", ErrUserAlreadyExists
	}

	uuid, err := unique.UUID()
	if err != nil {
		return "", err
	}
	hashedPassword, err := hash.FromPassword([]byte(strings.TrimSpace(in.GetPassword())))
	if err != nil {
		return "", err
	}
	var newUser = types.UserAuthInfo{
		Uuid:         uuid,
		Username:     strings.TrimSpace(in.GetUsername()),
		Organization: strings.TrimSpace(in.GetOrganisation()),
		Password:     hashedPassword,
	}
	err = repo.InsertOne(ctx, config.UserAuthDB, config.UserAuthColl, newUser)
	if err != nil {
		return "", err
	}
	return uuid, nil
}
