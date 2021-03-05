package app

// why are they here?????? change this pls
const (
	// DB, Collection names
	dbName        = "datalabs_app"
	appCollection = "app"
)

// convert to interface app!

type App struct{}

// New returns a new APP -> change to App interface!
func New() App {
	return App{}
}

// AppItem represents one App in the database ? do we need this? don't we have a def in the grpc already???
type AppItem struct {
	// mongoDB pk (document key)
	UUID       string `bson:"_id"`
	AppName    string `bson:"appName"`
	OwnerUUID  string `bson:"ownerUUID"`
	OrgnDomain string `bson:"orgnDomain"`
	// Member is a list of user_uuids mapped to this app
	Member   []string `bson:"member"`
	AppToken string   `bson:"appToken"`
}
