package users

import "github.com/KonstantinGasser/datalab/common"

type RegisterRequest struct {
	UserUuid     string `json:",omitempty"`
	Username     string `json:"username" required:"yes"`
	FirstName    string `json:"firstname" required:"yes"`
	LastName     string `json:"lastname" required:"yes"`
	Organization string `json:"organization" required:"yes"`
	Position     string `json:"position" required:"yes"`
	Password     string `json:"password" required:"yes"`
}
type RegisterResponse struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"error,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username" required:"yes"`
	Password string `json:"password" required:"yes"`
}
type LoginResponse struct {
	Status      int32  `json:"status"`
	Msg         string `json:"msg"`
	Err         string `json:"error,omitempty"`
	AccessToken string `json:"access_token"`
}

type UpdateProfileRequest struct {
	UserUuid  string `json:",omitempty"`
	FirstName string `json:"firstname" required:"yes"`
	LastName  string `json:"lastname" required:"yes"`
	Position  string `json:"position" required:"yes"`
}
type UpdateProfileResponse struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"error,omitempty"`
}

type GetProfileRequest struct {
	UserUuid string
}
type GetProfileResponse struct {
	Status int32            `json:"status"`
	Msg    string           `json:"msg"`
	Err    string           `json:"error,omitempty"`
	User   *common.UserInfo `json:"user"`
}

type GetColleagueRequest struct {
	Organization string
}
type GetColleagueResponse struct {
	Status     int32              `json:"status"`
	Msg        string             `json:"msg"`
	Err        string             `json:"error,omitempty"`
	Colleagues []*common.UserInfo `json:"user"`
}

type GetInvitableUsersRequest struct {
	AppUuid      string `json:"app_uuid"`
	Organization string
	UserUuids    []string
	AppMember    []*common.AppMember
}
type GetInvitableUsersResponse struct {
	Status    int32           `json:"status"`
	Msg       string          `json:"msg"`
	Err       string          `json:"error,omitempty"`
	Invitable []InvitableUser `json:"user"`
}

type InvitableUser struct {
	Uuid         string `json:"uuid"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Organization string `json:"orgn"`
	Position     string `json:"position"`
	Status       int32  `json:"status"`
}
