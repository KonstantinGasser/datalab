package types

type AppRole int32

const (
	Owner = iota
	Edit
	Viewer
)

type Permissions struct {
	UserUuid string          `bson:"_id"`
	UserOrgn string          `bson:"user_orgn"`
	Apps     []AppPermission `bson:"apps"`
}

type AppPermission struct {
	AppUuid string  `bson:"uuid"`
	Role    AppRole `bson:"role"`
}
