package update

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.user-administer/config"
	"github.com/KonstantinGasser/datalab/service.user-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/service.user-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

func User(ctx context.Context, repo repo.Repo, in *proto.UpdateRequest) error {

	var user = in.GetUser()

	var updateableInfo = types.UserUpdateable{
		Uuid:          in.GetCallerUuid(),
		FirstName:     user.GetFirstName(),
		LastName:      user.GetLastName(),
		OrgnPosition:  user.GetOrgnPosition(),
		ProfileImgURL: user.GetProfileImgUrl(),
	}
	query := bson.D{
		{
			Key:   "$set",
			Value: updateableInfo,
		},
	}
	_, err := repo.UpdateOne(ctx, config.UserDB, config.UserColl, bson.M{"_id": in.GetCallerUuid()}, query, false)
	if err != nil {
		return err
	}
	return nil
}
