package users

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
	Stauts int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"error,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username" required:"yes"`
	Password string `json:"password" required:"yes"`
}
type LoginResponse struct {
	Stauts      int32  `json:"status"`
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
	Stauts int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"error,omitempty"`
}
