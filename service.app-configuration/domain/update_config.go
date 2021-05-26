package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-configuration/errors"
)

const (
	UpdateFlagFunnel   = "funnel"
	UpdateFlagCampaign = "campaign"
	UpdateFlagBtnTime  = "btn_time"
)

var (
	ErrInvalidFlag     = fmt.Errorf("provided update flag is invalid")
	ErrUpdateOfNilData = fmt.Errorf("provided data would replace current data with nil")
)

func (svc appconfig) updateConfigByFlag(ctx context.Context, flag string, uuid string, config []interface{}) errors.ErrApi {

	switch flag {
	case UpdateFlagFunnel, UpdateFlagCampaign, UpdateFlagBtnTime:
		break
	default:
		return errors.New(http.StatusBadRequest, ErrInvalidFlag, "Provided flag is invalid")
	}

	err := svc.dao.UpdateByFlag(ctx, flag, uuid, config)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not update App Config")
	}
	return nil
}
