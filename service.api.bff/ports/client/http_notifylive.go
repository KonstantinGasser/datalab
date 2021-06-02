package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type VueMutation string

const (
	MutationAppInvite      VueMutation = "APP_INVITE"
	MutationInviteReminder VueMutation = "APP_INVITE_REMINDER"
	MutationSyncApp        VueMutation = "SYNC_APP"
)

type ClientNotifiyLive struct {
	addr       string
	apiPublish string
	apiRemove  string
}

func NewClientNotifyLive(addr string) *ClientNotifiyLive {
	return &ClientNotifiyLive{
		addr:       addr,
		apiPublish: "/api/v1/datalab/publish/event",
		apiRemove:  "/api/v1/datalab/remove/event",
	}
}

func (client ClientNotifiyLive) EmitSendEvent(ctx context.Context, event int, mutation VueMutation, receiverUuid, receiverOrgn string, value map[string]interface{}) error {

	var payload = map[string]interface{}{
		"receiver_uuid": receiverUuid,
		"receiver_orgn": receiverOrgn,
		"timestamp":     time.Now().Unix(),
		"mutation":      mutation,
		"event":         event,
		"value":         value,
	}
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(payload)

	req, _ := http.NewRequest("POST", client.AddrPublish(), buf)
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		logrus.Warnf("[client.Notify.EmitSendEvent] could not send invite: %v\n", err)
	}
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()
	return nil
}

func (client ClientNotifiyLive) EmitSendRemove(ctx context.Context, userUuid string, timestamp int64) error {

	var payload = map[string]interface{}{
		"user_uuid": userUuid,
		"timestamp": timestamp,
	}
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(payload)

	req, _ := http.NewRequest("POST", client.AddrRemove(), buf)
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		logrus.Warnf("[client.Notify.EmitSendRemove] could not send invite: %v\n", err)
	}
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()
	return nil
}

func (client ClientNotifiyLive) AddrPublish() string {
	return fmt.Sprintf("http://%s%s", client.addr, client.apiPublish)
}

func (client ClientNotifiyLive) AddrRemove() string {
	return fmt.Sprintf("http://%s%s", client.addr, client.apiRemove)
}
