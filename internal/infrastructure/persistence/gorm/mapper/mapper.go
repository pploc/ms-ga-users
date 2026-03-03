package mapper

import (
	"ms-ga-user/internal/domain/entity"
	"ms-ga-user/internal/infrastructure/persistence/gorm/model"
)

func ToEntityUser(u *model.User) *entity.User {
	if u == nil {
		return nil
	}
	return &entity.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
		AvatarURL: u.AvatarURL,
		Status:    entity.UserStatus(u.Status),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToModelUser(u *entity.User) *model.User {
	if u == nil {
		return nil
	}
	return &model.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
		AvatarURL: u.AvatarURL,
		Status:    string(u.Status),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToEntityUserProfile(p *model.UserProfile) *entity.UserProfile {
	if p == nil {
		return nil
	}
	return &entity.UserProfile{
		UserID:                p.UserID,
		Department:            p.Department,
		HireDate:              p.HireDate,
		EmergencyContactName:  p.EmergencyContactName,
		EmergencyContactPhone: p.EmergencyContactPhone,
		Notes:                 p.Notes,
		UpdatedAt:             p.UpdatedAt,
	}
}

func ToModelUserProfile(p *entity.UserProfile) *model.UserProfile {
	if p == nil {
		return nil
	}
	return &model.UserProfile{
		UserID:                p.UserID,
		Department:            p.Department,
		HireDate:              p.HireDate,
		EmergencyContactName:  p.EmergencyContactName,
		EmergencyContactPhone: p.EmergencyContactPhone,
		Notes:                 p.Notes,
		UpdatedAt:             p.UpdatedAt,
	}
}

func ToEntityUserAddress(a *model.UserAddress) *entity.UserAddress {
	if a == nil {
		return nil
	}
	return &entity.UserAddress{
		ID:        a.ID,
		UserID:    a.UserID,
		Street:    a.Street,
		City:      a.City,
		State:     a.State,
		ZipCode:   a.ZipCode,
		Country:   a.Country,
		IsPrimary: a.IsPrimary,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func ToModelUserAddress(a *entity.UserAddress) *model.UserAddress {
	if a == nil {
		return nil
	}
	return &model.UserAddress{
		ID:        a.ID,
		UserID:    a.UserID,
		Street:    a.Street,
		City:      a.City,
		State:     a.State,
		ZipCode:   a.ZipCode,
		Country:   a.Country,
		IsPrimary: a.IsPrimary,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
