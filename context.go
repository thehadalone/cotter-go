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

// SetUserID returns a copy of the context with the provided Cotter user ID.
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}
