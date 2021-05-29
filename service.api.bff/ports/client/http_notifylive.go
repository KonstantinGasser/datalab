package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type ClientNotifiyLive struct {
	addr       string
	apiPublish string
}

func NewClientNotifyLive(addr string) *ClientNotifiyLive {
	return &ClientNotifiyLive{
		addr:       addr,
		apiPublish: "/api/v1/datalab/publish/event",
	}
}

func (client ClientNotifiyLive) EmitSendInvite(ctx context.Context, receiverUuid, receiverOrgn string, value map[string]interface{}) error {

	var payload = map[string]interface{}{
		"receiver_uuid": receiverUuid,
		"receiver_orgn": receiverOrgn,
		"timestamp":     time.Now().Unix(),
		"mutation":      "APP_INVITE",
		"event":         0,
		"value":         value,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	_, err = http.NewRequest(http.MethodPost, client.Addr(), bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	// if resp.Response.StatusCode != http.StatusOK {
	// 	logrus.Errorf("[notifyService.EmitSendInvite]: %v\n", resp)
	// }
	return nil
}

func (client ClientNotifiyLive) Addr() string {
	return "http://localhost:8008/api/v1/datalab/publish/event"
	// return fmt.Sprintf("http://%s%s", client.addr, client.apiPublish)
}
