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
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	modelUser := mapper.ToModelUser(user)
	if err := r.db.WithContext(ctx).Create(modelUser).Error; err != nil {
		return nil, err
	}
	return mapper.ToEntityUser(modelUser), nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var modelUser model.User
	if err := r.db.WithContext(ctx).First(&modelUser, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if not found
		}
		return nil, err
	}
	return mapper.ToEntityUser(&modelUser), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var modelUser model.User
	if err := r.db.WithContext(ctx).First(&modelUser, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.ToEntityUser(&modelUser), nil
}

func (r *userRepository) GetAll(ctx context.Context, filter repository.UserFilter) ([]*entity.User, int64, error) {
	var modelUsers []model.User
	var total int64

	query := r.db.WithContext(ctx).Model(&model.User{})

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.Search != "" {
		searchLike := "%" + filter.Search + "%"
		query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ?", searchLike, searchLike, searchLike)
	}

	query.Count(&total)

	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
		if filter.Page > 0 {
			query = query.Offset((filter.Page - 1) * filter.Limit)
		}
	}

	if err := query.Order("created_at DESC").Find(&modelUsers).Error; err != nil {
		return nil, 0, err
	}

	var users []*entity.User
	for i := range modelUsers {
		users = append(users, mapper.ToEntityUser(&modelUsers[i]))
	}

	return users, total, nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	modelUser := mapper.ToModelUser(user)
	if err := r.db.WithContext(ctx).Model(modelUser).Updates(modelUser).Error; err != nil {
		return nil, err
	}
	return mapper.ToEntityUser(modelUser), nil
}

func (r *userRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status entity.UserStatus) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("status", string(status)).Error
}

func (r *userRepository) SoftDelete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, "id = ?", id).Error
}
