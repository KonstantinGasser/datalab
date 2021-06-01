package apps

import "github.com/KonstantinGasser/datalab/common"

type CreateAppRequest struct {
	AppName      string `json:"app_name" required:"yes"`
	Organization string `json:"organization"`
	AppUrl       string `json:"app_url" required:"yes"`
	AppDesc      string `json:"app_desc" required:"yes"`
	OwnerUuid    string
}
type CreateAppResponse struct {
	Status  int32  `json:"status"`
	Msg     string `json:"msg"`
	Err     string `json:"error,omitempty"`
	AppUuid string `json:"app_uuid"`
}

type GetAppRequest struct {
	AppUuid    string
	AuthedUser *common.AuthedUser
}
type GetAppResponse struct {
	Status int32                     `json:"status"`
	Msg    string                    `json:"msg"`
	Err    string                    `json:"error,omitempty"`
	App    *common.AppInfo           `json:"app"`
	Config *common.AppConfigurations `json:"config"`
	Token  *common.AppAccessToken    `json:"token"`
	Owner  *common.UserInfo          `json:"owner"`
}

type GetAppListRequest struct {
	AuthedUser *common.AuthedUser
}
type GetAppListResponse struct {
	Status int32               `json:"status"`
	Msg    string              `json:"msg"`
	Err    string              `json:"error,omitempty"`
	Apps   []*common.AppSubset `json:"apps"`
}

type CreateAppTokenRequest struct {
	AuthedUser *common.AuthedUser
	AppUuid    string `json:"app_uuid"`
	// these are inputs from the user which the user must provid
	// in order to verify the issuing of an app token
	AppName      string `json:"app_name"`
	Organization string `json:"orgn_domain"`
}
type CreateAppTokenResponse struct {
	Status   int32                  `json:"status"`
	Msg      string                 `json:"msg"`
	Err      string                 `json:"error,omitempty"`
	AppToken *common.AppAccessToken `json:"app_token"`
}

type UpdateConfigRequest struct {
	AuthedUser *common.AuthedUser
	UpdateFlag string           `json:"flag"`
	AppRefUuid string           `json:"app_uuid"`
	Stages     []*common.Stage  `json:"stages"`
	Records    []*common.Record `json:"records"`
	BtnDefs    []*common.BtnDef `json:"btn_defs"`
}
type UpdateConfigResponse struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"error,omitempty"`
}

type SendInviteRequest struct {
	AuthedUser  *common.AuthedUser
	InvitedUuid string `json:"invited_uuid"`
	AppUuid     string `json:"app_uuid"`
}
type SendInviteResponse struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"error,omitempty"`
}

type AcceptInviteRequest struct {
	AuthedUser            *common.AuthedUser
	AppUuid               string `json:"app_uuid"`
	NotificationTimestamp int64  `json:"event_timestamp"`
}
type AcceptInviteResponse struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"error,omitempty"`
}
