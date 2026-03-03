package model

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	UserID                uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Department            *string    `gorm:"type:varchar(100)"`
	HireDate              *time.Time `gorm:"type:date"`
	EmergencyContactName  *string    `gorm:"type:varchar(200)"`
	EmergencyContactPhone *string    `gorm:"type:varchar(20)"`
	Notes                 *string    `gorm:"type:text"`
	UpdatedAt             time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	User                  User       `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
