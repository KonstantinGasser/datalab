package util

import "context"

// StringValueCtx takes a context and a key and returns the value of the key
// or the default value of string ""
func StringValueCtx(ctx context.Context, key string) string {
	value := ctx.Value(key)
	if value == nil {
		return ""
	}
	return value.(string)
}
