package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	FirstName string         `gorm:"type:varchar(100);not null"`
	LastName  string         `gorm:"type:varchar(100);not null"`
	Email     string         `gorm:"type:varchar(255);not null;uniqueIndex"`
	Phone     *string        `gorm:"type:varchar(20)"`
	AvatarURL *string        `gorm:"type:text"`
	Status    string         `gorm:"type:varchar(20);not null;default:'active';index"`
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
