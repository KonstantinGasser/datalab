package types

type AppInfo struct {
	// mongoDB pk (document key)
	Uuid        string   `bson:"_id" required:"yes"`
	AppName     string   `bson:"name" required:"yes"`
	URL         string   `bson:"url" required:"yes"`
	OwnerUuid   string   `bson:"owner_uuid" required:"yes"`
	OrgnDomain  string   `bson:"orgn_domain" required:"yes"`
	Description string   `bson:"description"`
	Member      []string `bson:"member"`
	// AppToken    string   `bson:"app_token"`
	// ConfigRef   string   `bson:"config_ref" required:"yes"`
	AppHash string `bson:"orgn_and_app_hash"`
}

type AppMetaInfo struct {
	Uuid    string `bson:"_id" required:"yes"`
	AppName string `bson:"name" required:"yes"`
}
