package domain

import "time"

// AppToken represents the token data as it will be stored in the datbase
type AppToken struct {
	AppUuid   string    `bson:"_id" required:"yes"`
	AppHash   string    `bson:"app_hash" required:"yes"`
	AppOwner  string    `bson:"app_owner" required:"yes"`
	AppOrigin string    `bson:"app_origin"`
	AppToken  string    `bson:"app_token"`
	Exp       time.Time `bson:"token_exp"`
}
