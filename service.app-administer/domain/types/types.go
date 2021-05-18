package types

type InviteStatus int

const (
	InvitePending InviteStatus = iota
	InviteAccepted
	InviteRejected
)

type AppInfo struct {
	// mongoDB pk (document key)
	Uuid        string    `bson:"_id" required:"yes"`
	AppName     string    `bson:"name" required:"yes"`
	URL         string    `bson:"url" required:"yes"`
	OwnerUuid   string    `bson:"owner_uuid" required:"yes"`
	OrgnDomain  string    `bson:"orgn_domain" required:"yes"`
	Description string    `bson:"description"`
	Invites     []Invites `bson:"member"`
	AppHash     string    `bson:"app_hash"`
}

type AppMetaInfo struct {
	Uuid    string `bson:"_id" required:"yes"`
	AppName string `bson:"name" required:"yes"`
}

type Invites struct {
	Uuid   string       `bson:"uuid"`
	Status InviteStatus `bson:"status"`
}
