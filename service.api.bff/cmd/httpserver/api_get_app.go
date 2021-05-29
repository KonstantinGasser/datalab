package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/sirupsen/logrus"
)

func (s Server) GetApp(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("[%v][Server.GetApp] received request: %v\n", tracingID, r.Host)

	authedUser := ctx_value.GetAuthedUser(r.Context())

	appUuid := r.URL.Query().Get("app")
	var request = apps.GetAppRequest{
		AppUuid:    appUuid,
		AuthedUser: authedUser,
	}

	resp := s.appCollectService.GetApp(r.Context(), &request)
	if resp.Status != http.StatusOK {
		logrus.Errorf("[%v][Server.GetApp] could not create app: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Status, resp)
}

func (s Server) GetAppList(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("[%v][Server.GetAppList] received request: %v\n", tracingID, r.Host)

	authedUser := ctx_value.GetAuthedUser(r.Context())

	resp := s.appCollectService.GetAppList(r.Context(), &apps.GetAppListRequest{AuthedUser: authedUser})
	if resp.Status != http.StatusOK {
		logrus.Errorf("[%v][Server.GetAppList] could not get apps: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Status, resp)
}
