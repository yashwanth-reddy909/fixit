package repo

import (
	"context"

	"fixit.com/backend/src/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
