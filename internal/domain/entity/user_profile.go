package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	UserID                uuid.UUID
	Department            *string
	HireDate              *time.Time
	EmergencyContactName  *string
	EmergencyContactPhone *string
	Notes                 *string
	UpdatedAt             time.Time
}
