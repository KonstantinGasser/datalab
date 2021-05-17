package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-configuration/errors"
	appconfigsc "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/service.eventmanager-live/domain/types"
)

type EventLogic interface {
	InitSession(ctx context.Context, session types.SessionStart) (*common.AppConfigInfo, errors.ErrApi)
}

type eventlogic struct {
	appConfigSvc appconfigsc.AppConfigurationClient
}

func NewEventLogic(appConfigSvc appconfigsc.AppConfigurationClient) EventLogic {
	return &eventlogic{
		appConfigSvc: appConfigSvc,
	}
}

func (svc eventlogic) InitSession(ctx context.Context, session types.SessionStart) (*common.AppConfigInfo, errors.ErrApi) {

	resp, err := svc.appConfigSvc.Get(ctx, &appconfigsc.GetRequest{
		Tracing_ID: "1",
		ForUuid:    session.AppUuid,
	})
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "could not load config information",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("%s", resp.GetMsg()),
		}
	}
	return resp.GetConfigs(), nil
}
