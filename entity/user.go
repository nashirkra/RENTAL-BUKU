package entity

import (
	"time"

	"gorm.io/gorm"
)

/**
 * This is a User entity class and is not intended for modification.
 */
type User struct {
	ID              uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name            string `gorm:"type:varchar(255)" json:"name"`
	Address         string `gorm:"type:text" json:"address"`
	Photo           string `gorm:"type:text" json:"photo"`
	Email           string `gorm:"uniqueIndex:varchar(255)" json:"email"`
	EmailVerifiedAt time.Time
	Password        string `gorm:"->;<-;not null" json:"-"`
	Role            string `gorm:"type:ENUM('admin','member');not null" json:"role"`
	DeletedAt       gorm.DeletedAt
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Token           string
}
