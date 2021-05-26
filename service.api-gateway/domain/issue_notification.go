package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
)

type Notification struct {
	ReceiverUuid string                 `json:"receiver_uuid"`
	ReceiverOrgn string                 `json:"receiver_orgn"`
	Mutation     string                 `json:"mutation"`
	Event        int                    `json:"event"`
	Value        map[string]interface{} `json:"value"`
}

type NotfyMsg struct {
	Mutation string                 `json:"mutation"`
	Event    int                    `json:"event"`
	Value    map[string]interface{} `json:"value"`
}

func (svc gatewaylogic) IssueNotification(ctx context.Context, notification Notification) errors.ErrApi {

	b, err := json.Marshal(&notification)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not send notification",
			Err:    err,
		}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8008/api/v1/datalab/publish/event", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	_, err = client.Do(req)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not send notification",
			Err:    err,
		}
	}
	return nil
}
