package repository

// MongoUser represents a User Object living
// in the mongo database
type MongoUser struct {
	UUID       string `bson:"_id"`
	Username   string `bson:"username"`
	Password   string `bson:"password"`
	OrgnDomain string `bson:"orgnDomain"`
}
