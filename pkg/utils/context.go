package utils

import (
	"MYSHOP/pkg/models"
	"context"
)

type contextKey string

const userIDKey contextKey = "userID"

func SetUserIDToContext(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserIDFromContext(ctx context.Context) *models.User {
	if userID, ok := ctx.Value(userIDKey).(uint); ok {
		userModel, _ := models.GetUserById(int64(userID))
		return userModel
	}
	panic("User id not found")
}
