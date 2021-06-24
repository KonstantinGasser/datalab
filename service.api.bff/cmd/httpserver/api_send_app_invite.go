package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/sirupsen/logrus"
)

func (s Server) SendAppInvite(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("[%v][Server.SendAppInvite] received request: %v\n", tracingID, r.Host)

	authedUser := ctx_value.GetAuthedUser(r.Context())

	var request apps.SendInviteRequest
	if err := s.decode(r.Body, &request); err != nil {
		s.onErr(w, http.StatusBadRequest, "Could not decode r.Body")
		return
	}
	request.AuthedUser = authedUser
	resp := s.appInviteService.SendInvite(r.Context(), &request)
	if resp.Status != http.StatusOK {
		logrus.Errorf("[%v][Server.SendAppInvite] could not send app invite: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Status, resp)
}

func (s Server) SendAppInviteReminder(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("[%v][Server.SendAppInviteReminder] received request: %v\n", tracingID, r.Host)

	authedUser := ctx_value.GetAuthedUser(r.Context())
	var request apps.InviteReminderRequest
	if err := s.decode(r.Body, &request); err != nil {
		s.onErr(w, http.StatusBadRequest, "Could not decode r.Body")
		return
	}
	request.AuthedUser = authedUser

	resp := s.appInviteService.SendReminder(r.Context(), &request)
	if resp.Status != http.StatusOK {
		logrus.Errorf("[%v][Server.SendAppInviteReminder] could not send invite reminder: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Status, resp)
}

func (s Server) AcceptInvite(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("[%v][Server.GetUserProfile] received request: %v\n", tracingID, r.Host)

	authedUser := ctx_value.GetAuthedUser(r.Context())

	var request apps.AcceptInviteRequest
	if err := s.decode(r.Body, &request); err != nil {
		s.onErr(w, http.StatusBadRequest, "Could not decode r.Body")
		return
	}
	request.AuthedUser = authedUser

	resp := s.appInviteService.AcceptInvite(r.Context(), &request)
	if resp.Status != http.StatusOK {
		logrus.Errorf("[%v][Server.AcceptInvite] could not accept invite: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Status, resp)
}
