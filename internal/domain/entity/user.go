package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus string

const (
	UserStatusActive     UserStatus = "active"
	UserStatusSuspended  UserStatus = "suspended"
	UserStatusTerminated UserStatus = "terminated"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Phone     *string
	AvatarURL *string
	Status    UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
