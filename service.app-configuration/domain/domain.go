package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-configuration/domain/permissions"
	"github.com/KonstantinGasser/datalab/service.app-configuration/errors"
	"github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppConfig interface {
	GetConfigs(ctx context.Context, in *proto.GetRequest) (*common.AppConfigInfo, errors.ErrApi)
	InitConfigs(ctx context.Context, in *proto.InitRequest) errors.ErrApi
	UpdateConfig(ctx context.Context, in *proto.UpdateRequest) errors.ErrApi
}

type appconfig struct {
	dao         Dao
	permissions permissions.Permission
}

func NewDomainLogic(dao Dao) AppConfig {
	return &appconfig{
		dao:         dao,
		permissions: permissions.New(dao),
	}
}

// initConfigs creates an initialized config struct which will be stored in the database
func (svc appconfig) InitConfigs(ctx context.Context, in *proto.InitRequest) errors.ErrApi {
	initConifg := ConfigInfo{
		AppUuid:  in.GetForApp(),
		Funnel:   make([]Stage, 0),
		Campaign: make([]Record, 0),
		BtnTime:  make([]BtnDef, 0),
	}
	return svc.initAppConfig(ctx, initConifg)
}

// GetConfigs looks up all configs maped to a given AppUuid if the caller has read access to the resource
func (svc appconfig) GetConfigs(ctx context.Context, in *proto.GetRequest) (*common.AppConfigInfo, errors.ErrApi) {
	permissionErr := svc.permissions.HasRead(ctx, in.GetAppUuid(), in.GetUserClaims().Permissions.GetApps())
	if permissionErr != nil {
		return nil, permissionErr
	}
	appConfig, err := svc.dao.GetById(ctx, in.GetAppUuid())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusNotFound, err, "Could not find App Config")
		}
		return nil, errors.New(http.StatusInternalServerError, err, "Could not get App Config")
	}
	// convert dao struct to API struct
	var funnel = make([]*common.Funnel, len(appConfig.Funnel))
	for i, item := range appConfig.Funnel {
		funnel[i] = &common.Funnel{Id: item.ID, Name: item.Name, Transition: item.Transition}
	}
	var campaign = make([]*common.Campaign, len(appConfig.Campaign))
	for i, item := range appConfig.Campaign {
		campaign[i] = &common.Campaign{Id: item.ID, Name: item.Name, Prefix: item.Suffix}
	}
	var btnTime = make([]*common.BtnTime, len(appConfig.BtnTime))
	for i, item := range appConfig.BtnTime {
		btnTime[i] = &common.BtnTime{Id: item.ID, Name: item.Name, BtnName: item.BtnName}
	}

	return &common.AppConfigInfo{
		Funnel:   funnel,
		Campaign: campaign,
		BtnTime:  btnTime,
	}, nil
}

func (svc appconfig) UpdateConfig(ctx context.Context, in *proto.UpdateRequest) errors.ErrApi {
	permissionErr := svc.permissions.HasRead(ctx, in.GetAppUuid(), in.GetUserClaims().Permissions.GetApps())
	if permissionErr != nil {
		return permissionErr
	}

	// translate from protobuf to dao struct
	var config []interface{}
	switch in.GetUpdateFlag() {
	case "funnel":
		config = make([]interface{}, len(in.GetStages()))
		for i, item := range in.GetStages() {
			config[i] = Stage{ID: item.Id, Name: item.Name, Transition: item.Transition}
		}
	case "campaign":
		config = make([]interface{}, len(in.GetRecords()))
		for i, item := range in.GetRecords() {
			config[i] = Record{ID: item.Id, Name: item.Name, Suffix: item.Prefix}
		}
	case "btn_time":
		config = make([]interface{}, len(in.GetBtnDefs()))
		for i, item := range in.GetBtnDefs() {
			config[i] = BtnDef{ID: item.Id, Name: item.Name, BtnName: item.BtnName}
		}
	default:
		return errors.New(http.StatusBadRequest, fmt.Errorf("could not match '%s' as update-flag", in.GetUpdateFlag()), "Provided update-flag not found")
	}

	return svc.updateConfigByFlag(ctx, in.GetUpdateFlag(), in.GetAppUuid(), config)
}
