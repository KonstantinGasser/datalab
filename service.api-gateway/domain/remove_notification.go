package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
)

type RemoveEvent struct {
	UserUuid string `json:"user_uuid"`
	Timesamp int64  `json:"timestamp"`
}

func (svc gatewaylogic) RemoveNotification(ctx context.Context, event RemoveEvent) error {
	b, err := json.Marshal(&event)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not send notification",
			Err:    err,
		}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8008/api/v1/datalab/remove/event", bytes.NewBuffer(b))
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
