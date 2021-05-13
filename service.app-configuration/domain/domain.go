package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-configuration/domain/get"
	"github.com/KonstantinGasser/datalab/service.app-configuration/domain/initialize"
	"github.com/KonstantinGasser/datalab/service.app-configuration/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-configuration/domain/update"
	"github.com/KonstantinGasser/datalab/service.app-configuration/errors"
	"github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/service.app-configuration/repo"
)

type AppConfig interface {
	GetConfigs(ctx context.Context, in *proto.GetRequest) (*common.AppConfigInfo, errors.ErrApi)
	InitConfigs(ctx context.Context, in *proto.InitRequest) errors.ErrApi
	UpdateConfig(ctx context.Context, in *proto.UpdateRequest) errors.ErrApi
}

type appconfig struct {
	repo repo.Repo
}

func NewAppConfigLogic(repo repo.Repo) AppConfig {
	return &appconfig{
		repo: repo,
	}
}

func (svc appconfig) InitConfigs(ctx context.Context, in *proto.InitRequest) errors.ErrApi {
	if err := initialize.Configs(ctx, svc.repo, in); err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not initialize configurations",
			Err:    err,
		}
	}
	return nil
}

func (svc appconfig) GetConfigs(ctx context.Context, in *proto.GetRequest) (*common.AppConfigInfo, errors.ErrApi) {
	cfgs, err := get.Configs(ctx, svc.repo, in)
	if err != nil {
		if err == get.ErrNotFound {
			return nil, errors.ErrAPI{
				Status: http.StatusNotFound,
				Msg:    "Could not find any related App-Configs",
				Err:    err,
			}
		}
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not find any related App-Configs",
			Err:    err,
		}
	}
	return cfgs, nil
}

func (svc appconfig) UpdateConfig(ctx context.Context, in *proto.UpdateRequest) errors.ErrApi {

	var cfg []types.Config
	switch in.GetUpdateFlag() {
	case "funnel":
		cfg = make([]types.Config, len(in.GetStages()))
		for i, item := range in.GetStages() {
			cfg[i] = types.Stage{ID: item.Id, Name: item.Name, Transition: item.Transition}
		}
	case "record":
		cfg = make([]types.Config, len(in.GetRecords()))
		for i, item := range in.GetRecords() {
			cfg[i] = types.Record{ID: item.Id, Name: item.Name, Prefix: item.Prefix}
		}
	case "btn_defs":
		cfg = make([]types.Config, len(in.GetBtnDefs()))
		for i, item := range in.GetBtnDefs() {
			cfg[i] = types.BtnDef{ID: item.Id, Name: item.Name, BtnName: item.BtnName}
		}
	default:
		return errors.ErrAPI{
			Status: http.StatusBadRequest,
			Msg:    "Provided update-flag not found",
			Err:    fmt.Errorf("could not match '%s' as update-flag", in.GetUpdateFlag()),
		}
	}

	err := update.ByFlag(ctx, svc.repo, in.GetUpdateFlag(), in.GetUUID(), cfg)
	if err != nil {
		if err == update.ErrInvalidFlag {
			return errors.ErrAPI{
				Status: http.StatusBadRequest,
				Msg:    "Provided update flag is invalid",
				Err:    err,
			}
		}
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not update config",
			Err:    err,
		}
	}
	return nil
}
