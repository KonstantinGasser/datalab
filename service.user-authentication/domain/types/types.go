package types

const (
	Owner = iota
	Edit
	Viewer
)

// UserAuthInfo represents the data how it is
// persisted in the database
type UserAuthInfo struct {
	Uuid         string `bson:"_id"`
	Username     string `bson:"username"`
	Organization string `bson:"organization"`
	Password     string `bson:"password"`
}

type AppRole int32

type Permissions struct {
	UserUuid string          `bson:"_id"`
	UserOrgn string          `bson:"user_orgn"`
	Apps     []AppPermission `bson:"apps"`
}

type AppPermission struct {
	AppUuid string  `bson:"uuid" json:"app_uuid"`
	Role    AppRole `bson:"role" json:"role"`
}
