package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
	"github.com/sirupsen/logrus"
)

func (s Server) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	tracingID := r.Context().Value("tracingID").(string)
	logrus.Infof("<%v>[Server.GetUserProfile] received request: %v\n", tracingID, r.Host)

	authedUser := r.Context().Value("user").(*common.AuthedUser)
	var request = &users.GetProfileRequest{
		UserUuid: authedUser.Uuid,
	}
	resp := s.userfetchService.FetchProfile(r.Context(), request)
	if resp.Stauts != http.StatusOK {
		logrus.Errorf("<%v>[Server.GetUserProfile] could not get user profile: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Stauts, resp)
}

func (s Server) GetColleagues(w http.ResponseWriter, r *http.Request) {
	tracingID := r.Context().Value("tracingID").(string)
	logrus.Infof("<%v>[Server.GetColleagues] received request: %v\n", tracingID, r.Host)

	authedUser := r.Context().Value("user").(*common.AuthedUser)
	var request = &users.GetColleagueRequest{
		UserUuid: authedUser.Uuid,
	}
	resp := s.userfetchService.FetchColleagues(r.Context(), request)
	if resp.Stauts != http.StatusOK {
		logrus.Errorf("<%v>[Server.GetColleagues] could not get colleague profiles: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Stauts, resp)
}
