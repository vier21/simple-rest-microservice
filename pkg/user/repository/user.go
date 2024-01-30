package repository

import (
	"context"
	"gorest/pkg/user/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	InsertOneUser(ctx context.Context, user domain.User) (*domain.User, error)
	InsertManyUser(ctx context.Context, users []domain.User) ([]domain.User, error)
	UpdateOneUser(ctx context.Context, user domain.User) (*domain.User, error)
	GetAllUser(ctx context.Context) ([]domain.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	DeleteOneUser(ctx context.Context, user domain.User) (*domain.User, error)
	DeleteManyUser(ctx context.Context, ids []uuid.UUID) (int, error)
}
