package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-configuration/errors"
	appconfigsc "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/service.eventmanager-live/domain/types"
	"github.com/KonstantinGasser/datalab/service.eventmanager-live/jwts"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type EventLogic interface {
	InitSession(ctx context.Context, session types.SessionStart) (*common.AppConfigInfo, string, errors.ErrApi)
	OpenSocket(ctx context.Context, ticket string, w http.ResponseWriter, r *http.Request) errors.ErrApi
}

type eventlogic struct {
	appConfigSvc appconfigsc.AppConfigurationClient
}

func NewEventLogic(appConfigSvc appconfigsc.AppConfigurationClient) EventLogic {
	return &eventlogic{
		appConfigSvc: appConfigSvc,
	}
}

func (svc eventlogic) InitSession(ctx context.Context, session types.SessionStart) (*common.AppConfigInfo, string, errors.ErrApi) {

	resp, err := svc.appConfigSvc.Get(ctx, &appconfigsc.GetRequest{
		Tracing_ID: "1",
		ForUuid:    session.AppUuid,
	})
	if err != nil {
		return nil, "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "could not load config information",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, "", errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("%s", resp.GetMsg()),
		}
	}
	ticket, err := jwts.WebSocketTicket(session.Cookie)
	if err != nil {
		return nil, "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "could not load config information",
			Err:    err,
		}
	}
	return resp.GetConfigs(), ticket, nil
}

func (svc eventlogic) OpenSocket(ctx context.Context, ticket string, w http.ResponseWriter, r *http.Request) errors.ErrApi {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not establish web-socket connection",
			Err:    err,
		}
	}
	fmt.Printf("Conn: %+v\n", conn)
	conn.WriteMessage(websocket.TextMessage, []byte("Hello World"))
	return nil
}
