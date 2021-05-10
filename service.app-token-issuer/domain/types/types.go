package types

import "time"

type AppToken struct {
	AppUuid   string    `bson:"_id" required:"yes"`
	AppHash   string    `bson:"app_hash" required:"yes"`
	AppOrigin string    `bson:"app_origin" required:"yes"`
	AppToken  string    `bson:"app_token"`
	Exp       time.Time `bson:"token_exp"`
}
