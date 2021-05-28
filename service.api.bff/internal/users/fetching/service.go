package fetching

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client"
)

type Service interface {
	FetchProfile(ctx context.Context, r *users.GetProfileRequest) *users.GetProfileResponse
	FetchColleagues(ctx context.Context, r *users.GetColleagueRequest) *users.GetColleagueResponse
}

type service struct {
	userMetaClient client.ClientUserMeta
}

func NewService(userMetaClient client.ClientUserMeta) Service {
	return &service{
		userMetaClient: userMetaClient,
	}
}

func (s service) FetchProfile(ctx context.Context, r *users.GetProfileRequest) *users.GetProfileResponse {

	user, err := s.userMetaClient.GetProfile(ctx, r)
	if err != nil {
		return &users.GetProfileResponse{
			Stauts: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
			User:   nil,
		}
	}
	return &users.GetProfileResponse{
		Stauts: http.StatusOK,
		Msg:    "User Profile",
		User:   user,
	}
}

func (s service) FetchColleagues(ctx context.Context, r *users.GetColleagueRequest) *users.GetColleagueResponse {

	colleagues, err := s.userMetaClient.GetColleagues(ctx, r)
	if err != nil {
		return &users.GetColleagueResponse{
			Stauts:     err.Code(),
			Msg:        err.Info(),
			Err:        err.Error(),
			Colleagues: nil,
		}
	}
	return &users.GetColleagueResponse{
		Stauts:     http.StatusOK,
		Msg:        "Colleagues",
		Colleagues: colleagues,
	}
}
