package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api-gateway/domain"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) AcceptInvite(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.AcceptInvite] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var inviteForm domain.AcceptInviteForm
	err := handler.decode(r.Body, &inviteForm)
	if err != nil {
		logrus.Errorf("<%v>[handler.AcceptInvite] could not decode r.Body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		handler.onError(w, "could not decode r.Body", http.StatusBadRequest)
		return
	}
	_, acceptErr := handler.domain.AcceptInvite(r.Context(), inviteForm)
	if err != nil {
		logrus.Errorf("<%v>[handler.AcceptInvite] could not accept user invite: %v\n", ctx_value.GetString(r.Context(), "tracingID"), acceptErr.Error())
		handler.onError(w, acceptErr.Info(), int(acceptErr.Code()))
		return
	}
	handler.onSuccessJSON(w, nil, http.StatusOK)
}
