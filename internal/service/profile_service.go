package service

import (
	"context"

	"ms-ga-user/internal/domain/entity"
	"ms-ga-user/internal/domain/repository"

	"github.com/google/uuid"
)

type ProfileService interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, []*entity.UserAddress, error)
	UpdateProfile(ctx context.Context, profile *entity.UserProfile, addresses []*entity.UserAddress) (*entity.UserProfile, error)
}

type profileServiceImpl struct {
	profileRepo repository.ProfileRepository
}

func NewProfileService(profileRepo repository.ProfileRepository) ProfileService {
	return &profileServiceImpl{
		profileRepo: profileRepo,
	}
}

func (s *profileServiceImpl) GetProfile(ctx context.Context, userID uuid.UUID) (*entity.UserProfile, []*entity.UserAddress, error) {
	profile, err := s.profileRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, nil, err
	}
	if profile == nil {
		profile = &entity.UserProfile{UserID: userID}
	}

	addresses, err := s.profileRepo.GetAddressesByUserID(ctx, userID)
	if err != nil {
		return nil, nil, err
	}

	return profile, addresses, nil
}

func (s *profileServiceImpl) UpdateProfile(ctx context.Context, profile *entity.UserProfile, addresses []*entity.UserAddress) (*entity.UserProfile, error) {
	updatedProfile, err := s.profileRepo.Upsert(ctx, profile)
	if err != nil {
		return nil, err
	}

	for _, addr := range addresses {
		addr.UserID = profile.UserID
		if _, err := s.profileRepo.UpsertAddress(ctx, addr); err != nil {
			return nil, err
		}
	}

	return updatedProfile, nil
}
