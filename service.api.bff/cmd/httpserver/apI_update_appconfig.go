package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/sirupsen/logrus"
)

func (s Server) UpdateAppConfig(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("<%v>[Server.UpdateAppConfig] received request: %v\n", tracingID, r.Host)

	authedUser := ctx_value.GetAuthedUser(r.Context())
	var request apps.UpdateConfigRequest
	if err := s.decode(r.Body, &request); err != nil {
		s.onErr(w, http.StatusBadRequest, "Could not decode r.Body")
		return
	}

	request.AuthedUser = authedUser
	resp := s.appModifyService.UpdateConfig(r.Context(), &request)
	if resp.Status != http.StatusOK {
		logrus.Errorf("<%v>[Server.UpdateAppConfig] could not update app config: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Status, resp)
}
