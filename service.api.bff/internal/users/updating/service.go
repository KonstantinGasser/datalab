package updating

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client"
	"github.com/KonstantinGasser/required"
)

type Service interface {
	UpdateProfile(ctx context.Context, r *users.UpdateProfileRequest) *users.UpdateProfileResponse
}

type service struct {
	userMetaClient client.ClientUserMeta
}

func NewService(userMetaClient client.ClientUserMeta) Service {
	return &service{
		userMetaClient: userMetaClient,
	}
}

func (s service) UpdateProfile(ctx context.Context, r *users.UpdateProfileRequest) *users.UpdateProfileResponse {
	if err := required.Atomic(r); err != nil {
		return &users.UpdateProfileResponse{
			Stauts: http.StatusBadRequest,
			Msg:    "Mandatory fields missing",
			Err:    err.Error(),
		}
	}
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return &users.UpdateProfileResponse{
			Stauts: http.StatusUnauthorized,
			Msg:    "User must be logged in",
			Err:    fmt.Errorf("user not logged in").Error(),
		}
	}
	r.UserUuuid = authedUser.Uuid

	err := s.userMetaClient.UpdateUserProfile(ctx, r)
	if err != nil {
		return &users.UpdateProfileResponse{
			Stauts: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}
	return &users.UpdateProfileResponse{
		Stauts: http.StatusOK,
		Msg:    "User Profile updated",
	}
}
