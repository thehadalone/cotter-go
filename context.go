package cotter

import "context"

type userIDKey struct{}

// UserID extracts Cotter user ID from the context.
func UserID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if id, ok := ctx.Value(userIDKey{}).(string); ok {
		return id
	}

	return ""
}
