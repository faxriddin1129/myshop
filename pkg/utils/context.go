package utils

import (
	"context"
)

type contextKey string

const userIDKey contextKey = "userID"

func SetUserIDToContext(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserIDFromContext(ctx context.Context) uint {
	if userID, ok := ctx.Value(userIDKey).(uint); ok {
		return userID
	}
	return 0
}
