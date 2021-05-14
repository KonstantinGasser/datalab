package get

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-configuration/config"
	"github.com/KonstantinGasser/datalab/service.app-configuration/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/service.app-configuration/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFound = fmt.Errorf("could not find any data")
)

func Configs(ctx context.Context, repo repo.Repo, in *proto.GetRequest) (*common.AppConfigInfo, error) {
	var cfg types.ConfigInfo
	err := repo.FindOne(ctx, config.CfgDB, config.CfgColl, bson.M{"_id": in.GetForUuid()}, &cfg)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var funnel = make([]*common.Funnel, len(cfg.Funnel))
	for i, item := range cfg.Funnel {
		funnel[i] = &common.Funnel{Id: item.ID, Name: item.Name, Transition: item.Transition}
	}
	var campaign = make([]*common.Campaign, len(cfg.Campaign))
	for i, item := range cfg.Campaign {
		campaign[i] = &common.Campaign{Id: item.ID, Name: item.Name, Prefix: item.Prefix}
	}
	var btnTime = make([]*common.BtnTime, len(cfg.BtnTime))
	for i, item := range cfg.BtnTime {
		btnTime[i] = &common.BtnTime{Id: item.ID, Name: item.Name, BtnName: item.BtnName}
	}

	return &common.AppConfigInfo{
		Funnel:   funnel,
		Campaign: campaign,
		BtnTime:  btnTime,
	}, nil
}
