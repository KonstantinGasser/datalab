package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
	"github.com/sirupsen/logrus"
)

func (s Server) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	tracingID := ctx_value.GetString(r.Context(), "tracingID")
	logrus.Infof("<%v>[Server.GetUserProfile] received request: %v\n", tracingID, r.Host)

	authedUser := ctx_value.GetAuthedUser(r.Context())
	var request = &users.GetProfileRequest{
		UserUuid: authedUser.Uuid,
	}
	resp := s.userfetchService.FetchProfile(r.Context(), request)
	if resp.Status != http.StatusOK {
		logrus.Errorf("<%v>[Server.GetUserProfile] could not get user profile: %v\n", tracingID, resp.Err)
	}
	s.onSuccess(w, resp.Status, resp)
}

// func (s Server) GetColleagues(w http.ResponseWriter, r *http.Request) {
// 	tracingID := ctx_value.GetString(r.Context(), "tracingID")
// 	logrus.Infof("<%v>[Server.GetColleagues] received request: %v\n", tracingID, r.Host)

// 	authedUser := ctx_value.GetAuthedUser(r.Context())
// 	var request = &users.GetColleagueRequest{
// 		UserUuid: authedUser.Uuid,
// 	}
// 	resp := s.userfetchService.FetchColleagues(r.Context(), request)
// 	if resp.Status != http.StatusOK {
// 		logrus.Errorf("<%v>[Server.GetColleagues] could not get colleague profiles: %v\n", tracingID, resp.Err)
// 	}
// 	s.onSuccess(w, resp.Status, resp)
// }
