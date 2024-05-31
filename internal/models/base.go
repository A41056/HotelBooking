package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

// Base model
type Base struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	ModifiedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"modified_at"`
}

func (base *Base) BeforeSave(db *gorm.DB) error {
	base.ModifiedAt = time.Now()
	return nil
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID for ID field
	user.ID = uuid.New()

	// Set CreatedAt to current time
	user.CreatedAt = time.Now()

	return nil
}
