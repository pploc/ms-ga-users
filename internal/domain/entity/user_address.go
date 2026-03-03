package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserAddress struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Street    string
	City      string
	State     *string
	ZipCode   *string
	Country   string
	IsPrimary bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
