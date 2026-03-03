package repository

import (
	"context"
	"errors"

	"ms-ga-user/internal/domain/entity"
	"ms-ga-user/internal/domain/repository"
	"ms-ga-user/internal/infrastructure/persistence/gorm/mapper"
	"ms-ga-user/internal/infrastructure/persistence/gorm/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) repository.ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, error) {
	var modelProfile model.UserProfile
	if err := r.db.WithContext(ctx).First(&modelProfile, "user_id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if not found
		}
		return nil, err
	}
	return mapper.ToEntityUserProfile(&modelProfile), nil
}

func (r *profileRepository) Upsert(ctx context.Context, profile *entity.UserProfile) (*entity.UserProfile, error) {
	modelProfile := mapper.ToModelUserProfile(profile)
	if err := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		UpdateAll: true,
	}).Create(modelProfile).Error; err != nil {
		return nil, err
	}
	return mapper.ToEntityUserProfile(modelProfile), nil
}

func (r *profileRepository) GetAddressesByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.UserAddress, error) {
	var modelAddresses []model.UserAddress
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&modelAddresses).Error; err != nil {
		return nil, err
	}

	var addresses []*entity.UserAddress
	for i := range modelAddresses {
		addresses = append(addresses, mapper.ToEntityUserAddress(&modelAddresses[i]))
	}
	return addresses, nil
}

func (r *profileRepository) UpsertAddress(ctx context.Context, address *entity.UserAddress) (*entity.UserAddress, error) {
	modelAddress := mapper.ToModelUserAddress(address)
	if err := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(modelAddress).Error; err != nil {
		return nil, err
	}
	return mapper.ToEntityUserAddress(modelAddress), nil
}
