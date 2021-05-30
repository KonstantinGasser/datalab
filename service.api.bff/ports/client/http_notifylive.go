package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
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
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(payload)

	req, _ := http.NewRequest("POST", client.Addr(), buf)
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		logrus.Warnf("[client.Notify.EmitSendInvite] could not send invite: %v\n", err)
	}
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()
	return nil
}

func (client ClientNotifiyLive) Addr() string {
	return "http://192.168.0.232:8008/api/v1/datalab/publish/event"
	// return fmt.Sprintf("http://%s%s", client.addr, client.apiPublish)
}
