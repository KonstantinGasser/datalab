package types

type UserAuthInfo struct {
	Uuid         string `bson:"_id"`
	Username     string `bson:"username"`
	Organization string `bson:"organization"`
	Password     string `bson:"password"`
}
