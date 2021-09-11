package auth

import (
	"github.com/google/uuid"
)

type UserClaims struct {
	UUID uuid.UUID `json:"uuid"`
}
