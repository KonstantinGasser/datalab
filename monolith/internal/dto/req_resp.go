package dto

import (
	"encoding/json"
	"net/http"

	"github.com/KonstantinGasser/required"
	"github.com/pkg/errors"
)

type ReqRegisterUser struct {
	UserUuid     string `json:",omitempty"`
	Username     string `json:"username" required:"yes"`
	FirstName    string `json:"firstname" required:"yes"`
	LastName     string `json:"lastname" required:"yes"`
	Organization string `json:"organization" required:"yes"`
	Position     string `json:"position" required:"yes"`
	Password     string `json:"password" required:"yes"`
}

type RespRegisterUser struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
	Err    string `json:"error,omitempty"`
}

func FromRequest(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	if err := required.Atomic(v); err != nil {
		return errors.Wrap(err, "missing values in request")
	}
	defer r.Body.Close()
	return nil
}
