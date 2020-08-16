package context

import (
	"context"

	"quacker/models"
)

type privateKey string

const userKey privateKey = "user"

// WithUser creates a new context, with User data added to it
func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// GetUser returns User data from the given context
func GetUser(ctx context.Context) *models.User {
	val := ctx.Value(userKey)
	if val != nil {
		user, ok := val.(*models.User) // 'ok' should be always true
		if ok {
			return user
		}
	}
	return nil
}
