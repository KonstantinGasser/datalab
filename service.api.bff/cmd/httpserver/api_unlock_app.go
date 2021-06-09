package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/sirupsen/logrus"
)

func (server Server) UnlockApp(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("[%v][Server.UnlockApp] received request: %v\n", tracingID, r.Host)

	var request apps.UnlockRequest
	if err := server.decode(r.Body, &request); err != nil {
		server.onErr(w, http.StatusBadRequest, "Could not decode r.Body")
		return
	}
	request.AuthedUser = ctx_value.GetAuthedUser(r.Context())
	resp := server.appModifyService.UnlockApp(r.Context(), &request)
	if resp.Status != http.StatusOK {
		logrus.Errorf("[%v][Server.UnlockApp] could not unlock app data: %v\n", tracingID, resp.Err)
	}
	server.onSuccess(w, resp.Status, resp)
}
