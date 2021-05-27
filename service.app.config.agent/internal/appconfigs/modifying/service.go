package modifying

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	UpdateFunnel(ctx context.Context, appRefUuid string, authedUser *common.AuthedUser, stages []appconfigs.Stage) errors.Api
	UpdateCampaign(ctx context.Context, appRefUuid string, authedUser *common.AuthedUser, records []appconfigs.Record) errors.Api
	UpdateBtnTime(ctx context.Context, appRefUuid string, authedUser *common.AuthedUser, btnDefs []appconfigs.BtnDef) errors.Api
}

type service struct {
	repo appconfigs.AppconfigRepo
}

func NewService(repo appconfigs.AppconfigRepo) Service {
	return &service{repo: repo}
}

// InitializeAppConfig creates the core data object to represent an AppConfig and stores it in the
// database
func (s *service) UpdateFunnel(ctx context.Context, appRefUuid string, authedUser *common.AuthedUser, stages []appconfigs.Stage) errors.Api {
	var storedAppConfig appconfigs.AppConfig
	err := s.repo.GetById(ctx, appRefUuid, &storedAppConfig)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New(http.StatusNotFound, err, "Could not find any App Config")
		}
		return errors.New(http.StatusInternalServerError, err, "Could not get App Config")
	}
	// check user permissions
	if err := storedAppConfig.HasReadOrWrite(authedUser.Uuid, authedUser.ReadWriteApps...); err != nil {
		return errors.New(http.StatusUnauthorized, err, "User has no rights to change App Config")
	}
	// update config
	storedAppConfig.ApplyFunnel(stages...)
	if err := s.repo.UpdateByFlag(ctx, appRefUuid, appconfigs.FlagFunnel, stages); err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not update Funnel Config")
	}
	return nil
}

func (s *service) UpdateCampaign(ctx context.Context, appRefUuid string, authedUser *common.AuthedUser, records []appconfigs.Record) errors.Api {
	var storedAppConfig appconfigs.AppConfig
	err := s.repo.GetById(ctx, appRefUuid, &storedAppConfig)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New(http.StatusNotFound, err, "Could not find any App Config")
		}
		return errors.New(http.StatusInternalServerError, err, "Could not get App Config")
	}
	// check user permissions
	if err := storedAppConfig.HasReadOrWrite(authedUser.Uuid, authedUser.ReadWriteApps...); err != nil {
		return errors.New(http.StatusUnauthorized, err, "User has no rights to change App Config")
	}
	// update config
	storedAppConfig.ApplyCampaign(records...)
	if err := s.repo.UpdateByFlag(ctx, appRefUuid, appconfigs.FlagCampaign, records); err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not update Campaign Config")
	}
	return nil
}

func (s *service) UpdateBtnTime(ctx context.Context, appRefUuid string, authedUser *common.AuthedUser, btnDefs []appconfigs.BtnDef) errors.Api {
	var storedAppConfig appconfigs.AppConfig
	err := s.repo.GetById(ctx, appRefUuid, &storedAppConfig)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New(http.StatusNotFound, err, "Could not find any App Config")
		}
		return errors.New(http.StatusInternalServerError, err, "Could not get App Config")
	}
	// check user permissions
	if err := storedAppConfig.HasReadOrWrite(authedUser.Uuid, authedUser.ReadWriteApps...); err != nil {
		return errors.New(http.StatusUnauthorized, err, "User has no rights to change App Config")
	}
	// update config
	storedAppConfig.ApplyBtnTime(btnDefs...)
	if err := s.repo.UpdateByFlag(ctx, appRefUuid, appconfigs.FlagBtnTime, btnDefs); err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not update Btn Time Config")
	}
	return nil
}
