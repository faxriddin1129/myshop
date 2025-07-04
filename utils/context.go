package utils

import (
	"context"
	"myshop/models"
)

type contextKey string

const userIDKey contextKey = "userID"

func SetUserIDToContext(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func Auth(ctx context.Context) *models.User {
	if userID, ok := ctx.Value(userIDKey).(uint); ok {
		userModel, _ := models.GetUserById(int64(userID))
		return userModel
	}
	panic("User id not found")
}
