package user

import "strings"

// DBUser represents a User Object living
// in the mongo database
// fields are not exported - setter must be used
// - this is done to avoid multiple strings.trim etc operation
type DBUser struct {
	UUID       string `bson:"_id"`
	Username   string `bson:"username"`
	Password   string `bson:"password"`
	OrgnDomain string `bson:"orgnDomain"`
}

// newDBUser returns a new DBUser but calls its setter to perform sanity checks
// on the input
func newDBUser(uuid, username, password, orgnDomain string) DBUser {
	u := DBUser{}
	u.setUUID(uuid)
	u.setUsername(username)
	u.setPassword(password)
	u.setOrgnDomain(orgnDomain)
	return u
}

// setUUID trims trailing spaces
func (u *DBUser) setUUID(uuid string) {
	u.UUID = strings.TrimSpace(uuid)
}

// setUsername trims trailing spaces
func (u *DBUser) setUsername(username string) {
	u.Username = strings.TrimSpace(username)
}

// setPassword trims trailing spaces
func (u *DBUser) setPassword(password string) {
	u.Password = strings.TrimSpace(password)
}

// setOrgnDomain trims trailing spaces
func (u *DBUser) setOrgnDomain(domain string) {
	u.OrgnDomain = strings.TrimSpace(domain)
}
