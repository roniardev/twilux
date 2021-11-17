package models

import (
	"time"

	"github.com/jkomyno/nanoid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	id, err := nanoid.Nanoid(10)
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("id", id)
	return nil
}
