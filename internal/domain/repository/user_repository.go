package repository

import (
	"context"

	"ms-ga-user/internal/domain/entity"

	"github.com/google/uuid"
)

type UserFilter struct {
	Status string
	Search string
	Page   int
	Limit  int
}

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetAll(ctx context.Context, filter UserFilter) ([]*entity.User, int64, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status entity.UserStatus) error
	SoftDelete(ctx context.Context, id uuid.UUID) error
}
