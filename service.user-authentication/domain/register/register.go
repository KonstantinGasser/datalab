package register

import (
	"context"
	"fmt"
	"strings"

	"github.com/KonstantinGasser/datalab/service.user-administer/repo"
	"github.com/KonstantinGasser/datalab/service.user-authentication/config"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/utils/unique"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrUserAlreadyExists = fmt.Errorf("user with this username already exists")
)

func NewUser(ctx context.Context, repo repo.Repo, in *proto.RegisterRequest) error {
	exists, err := repo.Exists(ctx, config.UserAuthDB, config.UserAuthColl, bson.M{"username": in.GetUsername()})
	if err != nil {
		return err
	}
	if exists {
		return ErrUserAlreadyExists
	}

	uuid, err := unique.UUID()
	if err != nil {
		return err
	}
	var newUser = types.UserAuthInfo{
		Uuid:         uuid,
		Username:     strings.TrimSpace(in.GetUsername()),
		Organization: strings.TrimSpace(in.GetOrganisation()),
		Password:     strings.TrimSpace(in.GetPassword()),
	}
	err = repo.InsertOne(ctx, config.UserAuthDB, config.UserAuthColl, newUser)
	if err != nil {
		return err
	}
	return nil
}
