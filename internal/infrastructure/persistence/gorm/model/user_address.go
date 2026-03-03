package model

import (
	"time"

	"github.com/google/uuid"
)

type UserAddress struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Street    string    `gorm:"type:varchar(255);not null"`
	City      string    `gorm:"type:varchar(100);not null"`
	State     *string   `gorm:"type:varchar(100)"`
	ZipCode   *string   `gorm:"type:varchar(20)"`
	Country   string    `gorm:"type:varchar(100);not null;default:'Thailand'"`
	IsPrimary bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
