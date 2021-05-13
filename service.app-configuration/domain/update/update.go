package update

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service.app-configuration/config"
	"github.com/KonstantinGasser/datalab/service.app-configuration/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-configuration/repo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrInvalidFlag = fmt.Errorf("provided update flag is invalid")
)

func ByFlag(ctx context.Context, repo repo.Repo, flag string, uuid string, cfg []types.Config) error {
	switch flag {
	case "funnel":
		break
	case "record":
		break
	case "btn_defs":
		break
	default:
		return ErrInvalidFlag
	}

	_, err := repo.UpdateOne(ctx, config.CfgDB, config.CfgColl, bson.M{"_id": uuid},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{{
					Key:   flag,
					Value: cfg,
				},
				},
			},
		}, false)
	return err
}
