package apps

import "github.com/KonstantinGasser/datalab/common"

type CreateAppRequest struct {
	AppName      string `json:"app_name" required:"yes"`
	Organization string `json:"organization" required:"yes"`
	AppUrl       string `json:"app_url" required:"yes"`
	AppDesc      string `json:"app_desc" required:"yes"`
	OwnerUuid    string
}
type CreateAppResponse struct {
	Stauts int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"error,omitempty"`
}

type GetAppRequest struct {
	AppUuid    string
	AuthedUser *common.AuthedUser
}
type GetAppResponse struct {
	Stauts int32                  `json:"status"`
	Msg    string                 `json:"msg"`
	Err    string                 `json:"error,omitempty"`
	App    *common.AppInfo        `json:"app"`
	Config *common.AppConfigInfo  `json:"config"`
	Token  *common.AppAccessToken `json:"token"`
	Owner  *common.UserInfo       `json:"owner"`
}
