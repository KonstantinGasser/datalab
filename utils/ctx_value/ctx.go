package ctx_value

import (
	"context"

	"github.com/KonstantinGasser/datalab/common"
)

type ctxKey string

// AddValue wrapps the context.WithValue function using a unified ctxKey to store
// the passed value. Returns a new context.Context with the value
func AddValue(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, ctxKey(key), value)
}

// GetString returns a string value from a given context or
// an empty string "" if value is nil
func GetString(ctx context.Context, key string) string {
	value := ctx.Value(ctxKey(key))
	if value == nil {
		return ""
	}
	return value.(string)
}

// GetAuthedUser is a special case where the API middleware adds the authed user
// coming from the TokenService to the context. Claims of user might be needed afterwards.
// GetAuthedUser is tide to the grpc_def of TokenService.AuthenticatedUser and if not found
// will return nil
func GetAuthedUser(ctx context.Context) *common.TokenClaims {
	value := ctx.Value(ctxKey("user"))
	if value == nil {
		return nil
	}
	return value.(*common.TokenClaims)
}
