package ctxkey

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/monolith/internal/domain/auth"
)

type (
	KeyString string
)

func Str(str string) KeyString {
	return KeyString(str)
}

func Session(ctx context.Context) (*auth.UserClaims, error) {
	value := ctx.Value(Str("authenticated-user"))
	if _, ok := value.(auth.UserClaims); !ok {
		return nil, fmt.Errorf("invalid user claims")
	}
	return value.(*auth.UserClaims), nil
}
