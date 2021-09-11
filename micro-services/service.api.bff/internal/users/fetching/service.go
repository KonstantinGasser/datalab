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
	FetchInvitableUsers(ctx context.Context, r *users.GetInvitableUsersRequest) *users.GetInvitableUsersResponse
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
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
			User:   nil,
		}
	}
	return &users.GetProfileResponse{
		Status: http.StatusOK,
		Msg:    "User Profile",
		User:   user,
	}
}

func (s service) FetchColleagues(ctx context.Context, r *users.GetColleagueRequest) *users.GetColleagueResponse {

	colleagues, err := s.userMetaClient.GetColleagues(ctx, r)
	if err != nil {
		return &users.GetColleagueResponse{
			Status:     err.Code(),
			Msg:        err.Info(),
			Err:        err.Error(),
			Colleagues: nil,
		}
	}
	return &users.GetColleagueResponse{
		Status:     http.StatusOK,
		Msg:        "User Profile",
		Colleagues: colleagues,
	}
}

func (s service) FetchInvitableUsers(ctx context.Context, r *users.GetInvitableUsersRequest) *users.GetInvitableUsersResponse {

	// get all users from the same organization
	userList, err := s.userMetaClient.GetColleagues(ctx, &users.GetColleagueRequest{
		Organization: r.Organization,
	})
	if err != nil {
		return &users.GetInvitableUsersResponse{
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}

	// need to merge them into a users.InvitableUser slice since the app.member.status is required
	// TODO: 2x for-loop looks not good. Might be a better way.
	var invitable = make([]users.InvitableUser, len(userList))
	for i, item := range userList {
		user := users.InvitableUser{
			Uuid:         item.Uuid,
			Username:     item.Username,
			FirstName:    item.FirstName,
			LastName:     item.LastName,
			Organization: item.OrgnDomain,
			Position:     item.OrgnPosition,
			Status:       0, // indicate no connection to app yet
			Avatar:       item.Avatar,
		}
		for _, member := range r.AppMember {
			if member.Uuid == user.Uuid {
				user.Status = member.Status
			}
		}
		invitable[i] = user
	}

	return &users.GetInvitableUsersResponse{
		Status:    http.StatusOK,
		Msg:       "Invitable users",
		Invitable: invitable,
	}
}
