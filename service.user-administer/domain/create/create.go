package create

import (
	"context"
	"fmt"
	"strings"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.user-administer/config"
	"github.com/KonstantinGasser/datalab/service.user-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/service.user-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrInvalidOrgnName = fmt.Errorf("organization name must not contain '/'")
	ErrUserNameTaken   = fmt.Errorf("username is already taken")
)

func User(ctx context.Context, repo repo.Repo, in *proto.CreateRequest) error {
	var user *common.UserInfo = in.GetUser()
	if ok := orgnNameAllowed(user.GetOrgnDomain()); !ok {
		return ErrInvalidOrgnName
	}

	taken, err := repo.Exists(ctx, config.UserDB, config.UserColl, bson.M{"username": user.GetUsername()})
	if err != nil {
		return err
	}
	if taken {
		return ErrUserNameTaken
	}

	var newUser = types.UserInfo{
		Uuid:          user.GetUuid(),
		Username:      strings.TrimSpace(user.GetUsername()),
		FirstName:     strings.TrimSpace(user.GetFirstName()),
		LastName:      strings.TrimSpace(user.GetLastName()),
		OrgnDomain:    strings.TrimSpace(user.GetOrgnDomain()),
		OrgnPosition:  strings.TrimSpace(user.GetOrgnPosition()),
		ProfileImgURL: "",
	}

	if err := repo.InsertOne(ctx, config.UserDB, config.UserColl, newUser); err != nil {
		return err
	}
	return nil
}
