package service

import (
	"context"

	"ms-ga-user/internal/domain/entity"
	"ms-ga-user/internal/domain/repository"
	"ms-ga-user/internal/infrastructure/messaging"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	ListUsers(ctx context.Context, filter repository.UserFilter) ([]*entity.User, int64, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeactivateUser(ctx context.Context, id uuid.UUID) error
}

type userServiceImpl struct {
	userRepo repository.UserRepository
	producer messaging.KafkaProducer
}

func NewUserService(userRepo repository.UserRepository, producer messaging.KafkaProducer) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
		producer: producer,
	}
}

func (s *userServiceImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	created, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	_ = s.producer.PublishEvent(messaging.EventUserCreated, map[string]interface{}{
		"user_id":    created.ID,
		"email":      created.Email,
		"first_name": created.FirstName,
		"last_name":  created.LastName,
	})

	return created, nil
}

func (s *userServiceImpl) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *userServiceImpl) ListUsers(ctx context.Context, filter repository.UserFilter) ([]*entity.User, int64, error) {
	return s.userRepo.GetAll(ctx, filter)
}

func (s *userServiceImpl) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	updated, err := s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	_ = s.producer.PublishEvent(messaging.EventUserUpdated, map[string]interface{}{
		"user_id":        updated.ID,
		"changed_fields": []string{"profile"}, // simplified for demonstration
	})

	return updated, nil
}

func (s *userServiceImpl) DeactivateUser(ctx context.Context, id uuid.UUID) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil || user == nil {
		return err // Return error or proceed depending on requirements
	}

	if err := s.userRepo.SoftDelete(ctx, id); err != nil {
		return err
	}
	if err := s.userRepo.UpdateStatus(ctx, id, entity.UserStatusTerminated); err != nil {
		return err
	}

	_ = s.producer.PublishEvent(messaging.EventUserDeactivated, map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
	})

	return nil
}
