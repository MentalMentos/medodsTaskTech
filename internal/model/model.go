package model

import "github.com/google/uuid"

type RefreshToken struct {
	ID          int       `gorm:"primaryKey"`
	UserID      uuid.UUID `gorm:"not null"`
	HashedToken string
	AccessJTI   string
	IPAddress   string
	Used        bool
}
