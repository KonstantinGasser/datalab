// put package in module to be available for all the services!!!!

package utils

import (
	"context"
)

// AddValCtx is a wrapper serving to unify the why of adding meta data to a given context.
// Calls context.WithValue(passed_context, key, value) returns new context.Context
func AddValCtx(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

// StringValueCtx takes a context and a key and returns the value of the key
// or the default value of string ""
func StringValueCtx(ctx context.Context, key string) string {
	value := ctx.Value(key)
	if value == nil {
		return ""
	}
	return value.(string)
}
