package update

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app-configuration/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-configuration/repo"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/config"
	"go.mongodb.org/mongo-driver/bson"
)

func ByFlag(ctx context.Context, repo repo.Repo, flag string, uuid string, cfg []types.Config) error {
	_, err := repo.UpdateOne(ctx, config.TokenDB, config.TokenColl, bson.M{"_id": uuid},
		bson.D{
			{Key: "$set",
				Value: bson.M{flag: cfg}},
		}, false)
	return err
}
