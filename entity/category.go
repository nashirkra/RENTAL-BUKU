package entity

import (
	"time"

	"gorm.io/gorm"
)

/**
 * This is a Category entity class and is not intended for modification.
 */
type Category struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
