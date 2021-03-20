package app

type app struct{}

// AppItem represents one App in the database ? do we need this? don't we have a def in the grpc already???
type AppItem struct {
	// mongoDB pk (document key)
	UUID        string `bson:"_id"`
	AppName     string `bson:"name"`
	OwnerUUID   string `bson:"owner_uuid"`
	OrgnDomain  string `bson:"orgn_domain"`
	Description string `bson:"description"`
	// Member is a list of user_uuids mapped to this app
	Member   []string `bson:"member"`
	Settings []string `bson:"setting"`
	AppToken string   `bson:"app_token"`
}

// AppItemLight is a minimum representation of an application
type AppItemLight struct {
	// mongoDB pk (document key)
	UUID    string `bson:"_id"`
	AppName string `bson:"name"`
}
