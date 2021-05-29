package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/sirupsen/logrus"
)

func (s Server) AcceptInvite(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("<%v>[Server.GetUserProfile] received request: %v\n", tracingID, r.Host)

	authedUser := ctx_value.GetAuthedUser(r.Context())

	var request apps.AcceptInviteRequest
	if err := s.decode(r.Body, &request); err != nil {
		s.onErr(w, http.StatusBadRequest, "Could not decode r.Body")
		return
	}
	request.AuthedUser = authedUser

	resp := s.appInviteService.AcceptInvite(r.Context(), &request)
	if resp.Status != http.StatusOK {
		logrus.Errorf("<%v>[Server.AcceptInvite] could not accept invite: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Status, resp)
}
