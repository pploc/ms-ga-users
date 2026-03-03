package handler

import (
	"ms-ga-user/internal/api/generated"
	"ms-ga-user/internal/domain/entity"
	"time"
)

func ToGeneratedUser(e *entity.User) *generated.User {
	if e == nil {
		return nil
	}
	id := e.ID.String()
	createdAt := e.CreatedAt.Format(time.RFC3339)
	updatedAt := e.UpdatedAt.Format(time.RFC3339)
	status := string(e.Status)

	return &generated.User{
		Id:        &id,
		FirstName: &e.FirstName,
		LastName:  &e.LastName,
		Email:     &e.Email,
		Phone:     e.Phone,
		AvatarUrl: e.AvatarURL,
		Status:    (*string)(&status),
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}
}

func ToGeneratedUserProfile(p *entity.UserProfile, addrs []*entity.UserAddress) *generated.UserProfile {
	if p == nil {
		return nil
	}
	userID := p.UserID.String()
	var hireDate *string
	if p.HireDate != nil {
		hd := p.HireDate.Format(time.DateOnly)
		hireDate = &hd
	}

	var generatedAddrs []generated.UserAddress
	for _, a := range addrs {
		id := a.ID.String()
		generatedAddrs = append(generatedAddrs, generated.UserAddress{
			Id:        &id,
			Street:    &a.Street,
			City:      &a.City,
			State:     a.State,
			ZipCode:   a.ZipCode,
			Country:   &a.Country,
			IsPrimary: &a.IsPrimary,
		})
	}

	return &generated.UserProfile{
		UserId:                &userID,
		Department:            p.Department,
		HireDate:              hireDate,
		EmergencyContactName:  p.EmergencyContactName,
		EmergencyContactPhone: p.EmergencyContactPhone,
		Addresses:             &generatedAddrs,
	}
}
