package hash

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	hashCost = 10
)

// FromPassword hashes the user's password before routed to other services (user-service)
// Only wraps the bcrypt.GenerateFromPassword func avoiding code duplication and a central point
// to change the cost of the hashing
func FromPassword(password []byte) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password can not be of length zero")
	}
	bytes, err := bcrypt.GenerateFromPassword(password, hashCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %v", err)
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a given password with one from the database
// Wraps the bcrypt.CompareHashAndPassword func to avoid code duplication and cleaner code
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
