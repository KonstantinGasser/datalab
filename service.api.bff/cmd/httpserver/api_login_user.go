package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
	"github.com/sirupsen/logrus"
)

func (s Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	tracingID := r.Context().Value("tracingID").(string)
	logrus.Infof("<%v>[Server.LoginUser] received request: %v\n", tracingID, r.Host)

	var request users.LoginRequest
	if err := s.decode(r.Body, &request); err != nil {
		s.onErr(w, http.StatusBadRequest, "Could not decode r.Body")
		return
	}
	resp := s.userauthService.Login(r.Context(), &request)
	if resp.Stauts != http.StatusOK {
		logrus.Errorf("<%v>[Server.LoginUser] could not login user: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Stauts, resp)
}
