package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
	"github.com/sirupsen/logrus"
)

func (s Server) RegisterUser(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("[%v][Server.RegisterUser] received request: %v\n", tracingID, r.Host)

	var request users.RegisterRequest
	if err := s.decode(r.Body, &request); err != nil {
		s.onErr(w, http.StatusBadRequest, "Could not decode r.Body")
		return
	}
	resp := s.userauthService.Register(r.Context(), &request)
	if resp.Status != http.StatusOK {
		logrus.Errorf("[%v][Server.RegisterUser] could not register user: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Status, resp)
}
