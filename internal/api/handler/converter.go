package handler

import (
	"ms-ga-user/internal/api/generated"
	"ms-ga-user/internal/domain/entity"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func ToGeneratedUser(e *entity.User) *generated.User {
	if e == nil {
		return nil
	}
	id := openapi_types.UUID(e.ID)
	email := openapi_types.Email(e.Email)
	status := generated.UserStatus(e.Status)
	createdAt := e.CreatedAt
	updatedAt := e.UpdatedAt

	return &generated.User{
		Id:        &id,
		FirstName: &e.FirstName,
		LastName:  &e.LastName,
		Email:     &email,
		Phone:     e.Phone,
		AvatarUrl: e.AvatarURL,
		Status:    &status,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}
}

func ToGeneratedUserProfile(p *entity.UserProfile, addrs []*entity.UserAddress) *generated.UserProfile {
	if p == nil {
		return nil
	}
	userID := openapi_types.UUID(p.UserID)
	var hireDate *openapi_types.Date
	if p.HireDate != nil {
		hd := openapi_types.Date{Time: *p.HireDate}
		hireDate = &hd
	}

	var generatedAddrs []generated.UserAddress
	for _, a := range addrs {
		addrID := openapi_types.UUID(a.ID)
		generatedAddrs = append(generatedAddrs, generated.UserAddress{
			Id:        &addrID,
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
