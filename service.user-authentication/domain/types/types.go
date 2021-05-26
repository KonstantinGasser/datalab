package types

// UserAuthInfo represents the User document in the database
type UserAuthInfo struct {
	Uuid         string `bson:"_id"`
	Username     string `bson:"username"`
	Organization string `bson:"organization"`
	Password     string `bson:"password"`
}