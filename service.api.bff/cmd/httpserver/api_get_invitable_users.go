package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
	"github.com/sirupsen/logrus"
)

func (s Server) GetInvitableUsers(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("[%v][Server.GetInvitableUsers] received request: %v\n", tracingID, r.Host)

	authedUser := ctx_value.GetAuthedUser(r.Context())

	var request users.GetInvitableUsersRequest
	if err := s.decode(r.Body, &request); err != nil {
		s.onErr(w, http.StatusBadRequest, "Could not decode r.Body")
		return
	}
	appRequest := apps.GetAppRequest{
		AppUuid:    request.AppUuid,
		AuthedUser: authedUser,
	}
	appResp := s.appCollectService.GetApp(r.Context(), &appRequest)
	if appResp.Status != http.StatusOK {
		s.onErr(w, appResp.Status, appResp.Msg)
		return
	}

	// convert app member to slice of uuids
	var uuids = make([]string, len(appResp.App.GetMember()))
	for i, item := range appResp.App.GetMember() {
		uuids[i] = item.GetUuid()
	}
	invitableRequest := users.GetInvitableUsersRequest{
		AppUuid:      request.AppUuid,
		UserUuids:    uuids,
		Organization: authedUser.Organization,
		AppMember:    appResp.App.GetMember(),
	}
	invitableResp := s.userfetchService.FetchInvitableUsers(r.Context(), &invitableRequest)
	if invitableResp.Status != http.StatusOK {
		logrus.Errorf("[%v][Server.GetInvitableUsers] could not get invitable users: %v\n", tracingID, invitableResp.Err)
	}
	s.onSuccess(w, invitableResp.Status, invitableResp)
}
