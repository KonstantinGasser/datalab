package types

type InviteStatus int

const (
	// InvitePending means the request has been send but not yet acknowledged
	InvitePending InviteStatus = iota + 1 // plus one else grpc will drop zero value
	// InviteAccepted means the requested user has acknowledged and accepted the invite
	InviteAccepted
	// InviteRejected means the requested user has acknowledged and rejected the invite
	InviteRejected
)

// AppMetaInfo represents a subset of the AppInfo only holding
// the app name and its uuid
type AppMetaInfo struct {
	Uuid    string `bson:"_id" required:"yes"`
	AppName string `bson:"name" required:"yes"`
}

// AppInfo represents the app data as it will be stored in the database
type AppInfo struct {
	// mongoDB pk (document key)
	Uuid        string   `bson:"_id" required:"yes"`
	AppName     string   `bson:"name" required:"yes"`
	URL         string   `bson:"url" required:"yes"`
	OwnerUuid   string   `bson:"owner_uuid" required:"yes"`
	OrgnDomain  string   `bson:"orgn_domain" required:"yes"`
	Description string   `bson:"description"`
	Invites     []Invite `bson:"member"`
	AppHash     string   `bson:"app_hash"`
}

// Invite refers to a user added to an app
type Invite struct {
	Uuid   string       `bson:"uuid"`
	Status InviteStatus `bson:"status"`
}
