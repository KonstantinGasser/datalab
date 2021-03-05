package util

import (
	"context"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
)

// StringValueCtx takes a context and a key and returns the value of the key
// or the default value of string ""
func StringValueCtx(ctx context.Context, key string) string {
	value := ctx.Value(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

// AtuhedUserValCtx returns an authenticated user struct as defined in the grpc_def.token_service.AuthenticatedUser
// should only be used to get the authenticated user claims from the WithAuth middleware
func AtuhedUserValCtx(ctx context.Context, key string) *tokenSrv.AuthenticatedUser {
	value := ctx.Value(key)
	if value == nil {
		return nil
	}
	return value.(*tokenSrv.AuthenticatedUser)
}
