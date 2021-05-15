package create

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.user-administer/config"
	"github.com/KonstantinGasser/datalab/service.user-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

func Compansate(ctx context.Context, repo repo.Repo, userUuid string) error {
	if err := repo.DeleteOne(ctx, config.UserDB, config.UserColl, bson.M{"_id": userUuid}); err != nil {
		return err
	}
	return nil
}
