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
	ErrInvalidFlag     = fmt.Errorf("provided update flag is invalid")
	ErrUpdateOfNilData = fmt.Errorf("provided data would replace current data with nil")
)

func ByFlag(ctx context.Context, repo repo.Repo, flag string, uuid string, cfg []types.Config) error {

	// check if the data matches the flag. Reason, if flag is set to one value but the passed data
	// describes a different config - the data behind the flag in the database will be overwritten with nil.
	// example: flag: funnel, types.ConfigInfo.Funnel = nil but types.ConfigInfo.Campaign = [{},{}]
	//		=> database record of config.funnel will be $set to nil since types.ConfigInfo.Funnel is nil
	switch flag {
	case "funnel":
		break
	case "campaign":
		break
	case "btn_time":
		break
	default:
		return ErrInvalidFlag
	}

	_, err := repo.UpdateOne(ctx, config.CfgDB, config.CfgColl, bson.M{"_id": uuid},
		bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{
						Key:   flag,
						Value: cfg,
					},
				},
			},
		}, false)
	return err
}
