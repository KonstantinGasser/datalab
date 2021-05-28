package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/sirupsen/logrus"
)

func (s Server) GetApp(w http.ResponseWriter, r *http.Request) {
	tracingID := r.Context().Value("tracingID").(string)
	logrus.Infof("<%v>[Server.GetApp] received request: %v\n", tracingID, r.Host)

	authedUser := r.Context().Value("user").(*common.AuthedUser)

	appUuid := r.URL.Query().Get("app")
	var request = apps.GetAppRequest{
		AppUuid:    appUuid,
		AuthedUser: authedUser,
	}

	resp := s.appCollectService.GetApp(r.Context(), &request)
	if resp.Stauts != http.StatusOK {
		logrus.Errorf("<%v>[Server.GetApp] could not create app: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Stauts, resp)
}
