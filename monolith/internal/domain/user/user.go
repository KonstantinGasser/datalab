package user

import (
	"strconv"
	"strings"
	"time"

	"github.com/KonstantinGasser/required"
)

type User struct {
	Uuid         string `bson:"_id"`
	Username     string `bson:"username" required:"yes" min:"5"`
	FirstName    string `bson:"first_name" required:"yes"`
	LastName     string `bson:"last_name" required:"yes"`
	Organization string `bson:"organization" required:"yes"`
	Position     string `bson:"position" required:"yes"`
	Avatar       string `bson:"avatar"`
}

var (
	// defautlAvatarApi allows to generare a random avatar based on some seed (like a timestamp etc)
	defaultAvatarApi = "https://avatars.dicebear.com/api/bottts/"
)

// UpdatableUser defins the fields of a User that can be changed
type UpdatableUser struct {
	Uuid                          string
	FirstName, LastName, Position string
}

func NewDefault(uuid, username, firstname, lastname, organization, position string) (*User, error) {
	// build random user avatar
	timeNow := strconv.FormatInt(time.Now().Unix(), 10)
	urlSeed := strings.Join([]string{timeNow, "svg"}, ".") // timetamp.svg
	avatarUrl := strings.Join([]string{defaultAvatarApi, urlSeed}, "/")
	user := &User{
		Uuid:         uuid,
		Username:     username,
		FirstName:    firstname,
		LastName:     lastname,
		Organization: organization,
		Position:     position,
		Avatar:       avatarUrl,
	}
	if err := required.Atomic(user); err != nil {
		return nil, err
	}
	return user, nil
}
