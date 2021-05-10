package types

type UserInfo struct {
	Uuid          string `bson:"_id" required:"yes"`
	Username      string `bson:"username" required:"yes"`
	FirstName     string `bson:"first_name" required:"yes"`
	LastName      string `bson:"last_name" required:"yes"`
	OrgnDomain    string `bson:"orgn_domain" required:"yes"`
	OrgnPosition  string `bson:"orgn_position" required:"yes"`
	ProfileImgURL string `bson:"profile_img_url" required:"yes"`
}

type UserMetaInfo struct {
	Uuid     string `bson:"_id" required:"yes"`
	Username string `bson:"username" required:"yes"`
}

type UserUpdateable struct {
	Uuid          string `bson:"_id" required:"yes"`
	FirstName     string `bson:"first_name" required:"yes"`
	LastName      string `bson:"last_name" required:"yes"`
	OrgnPosition  string `bson:"orgn_position" required:"yes"`
	ProfileImgURL string `bson:"profile_img_url"`
}
