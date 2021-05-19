package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api-gateway/domain"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) InviteUserToApp(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.InviteUserToApp] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var inviteForm domain.InviteForm
	err := handler.decode(r.Body, &inviteForm)
	if err != nil {
		logrus.Errorf("<%v>[handler.InviteUserToApp] could not decode r.Body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		handler.onError(w, "could not decode r.Body", http.StatusBadRequest)
		return
	}
	inviteErr := handler.domain.InviteToAppProcess(r.Context(), inviteForm)
	if err != nil {
		logrus.Errorf("<%v>[handler.InviteUserToApp] could not invite user: %v\n", ctx_value.GetString(r.Context(), "tracingID"), inviteErr.Error())
		handler.onError(w, inviteErr.Info(), int(inviteErr.Code()))
		return
	}
	handler.onSuccessJSON(w, nil, http.StatusOK)
}
