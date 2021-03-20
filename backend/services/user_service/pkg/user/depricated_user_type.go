package user

// import "strings"

// // DBUser represents a User Object living
// // in the mongo database
// // Fields must be exported in order for the struct be (un-)marshaled in/to bson
// type DBUser struct {
// 	UUID         string `bson:"_id"`
// 	Username     string `bson:"username"`
// 	Password     string `bson:"password"`
// 	FirstName    string `bson:"first_name"`
// 	LastName     string `bson:"last_name"`
// 	OrgnDomain   string `bson:"orgnDomain"`
// 	OrgnPosition string `bson:"orgn_position"`
// }

// // newDBUser returns a new DBUser but calls its setter to perform sanity checks
// // on the input - striping of white-spaces
// func newDBUser(uuid, username, firstName, lastName, password, orgnDomain, orgnPosition string) DBUser {
// 	u := DBUser{}
// 	u.setUUID(uuid)
// 	u.setUsername(username)
// 	u.setFirstName(firstName)
// 	u.setLastName(lastName)
// 	u.setPassword(password)
// 	u.setOrgnDomain(orgnDomain)
// 	u.setOrgnPosition(orgnPosition)
// 	return u
// }

// // setUUID trims trailing spaces
// func (u *DBUser) setUUID(uuid string) {
// 	u.UUID = strings.TrimSpace(uuid)
// }

// // setUsername trims trailing spaces
// func (u *DBUser) setUsername(username string) {
// 	u.Username = strings.TrimSpace(username)
// }

// // setFirstName trims trailing spaces
// func (u *DBUser) setFirstName(firstName string) {
// 	u.FirstName = strings.TrimSpace(firstName)
// }

// // setLastName trims trailing spaces
// func (u *DBUser) setLastName(lastName string) {
// 	u.LastName = strings.TrimSpace(lastName)
// }

// // setOrgnPosition trims trailing spaces
// func (u *DBUser) setOrgnPosition(orgnPosition string) {
// 	u.OrgnPosition = strings.TrimSpace(orgnPosition)
// }

// // setPassword trims trailing spaces
// func (u *DBUser) setPassword(password string) {
// 	u.Password = strings.TrimSpace(password)
// }

// // setOrgnDomain trims trailing spaces
// func (u *DBUser) setOrgnDomain(domain string) {
// 	u.OrgnDomain = strings.TrimSpace(domain)
// }
