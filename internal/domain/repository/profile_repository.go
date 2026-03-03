package repository

import (
	"context"

	"ms-ga-user/internal/domain/entity"

	"github.com/google/uuid"
)

type ProfileRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, error)
	Upsert(ctx context.Context, profile *entity.UserProfile) (*entity.UserProfile, error)
	GetAddressesByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserAddress, error)
	UpsertAddress(ctx context.Context, address *entity.UserAddress) (*entity.UserAddress, error)
}
