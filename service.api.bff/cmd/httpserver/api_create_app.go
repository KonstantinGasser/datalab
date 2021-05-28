package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/sirupsen/logrus"
)

func (s Server) CreateApp(w http.ResponseWriter, r *http.Request) {
	tracingID := r.Context().Value("tracingID").(string)
	logrus.Infof("<%v>[Server.CreateApp] received request: %v\n", tracingID, r.Host)

	authedUser := r.Context().Value("user").(*common.AuthedUser)

	var request apps.CreateAppRequest
	if err := s.decode(r.Body, &request); err != nil {
		s.onErr(w, http.StatusBadRequest, "Could not decode r.Body")
		return
	}
	request.OwnerUuid = authedUser.Uuid

	resp := s.appcreateService.CreateApp(r.Context(), &request)
	if resp.Stauts != http.StatusOK {
		logrus.Errorf("<%v>[Server.CreateApp] could not create app: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Stauts, resp)
}
