package entity

import (
	"time"

	"gorm.io/gorm"
)

/**
 * This is a User entity class and is not intended for modification.
 */
type User struct {
	ID              uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name            string    `gorm:"type:varchar(255);not null" json:"name"`
	Address         string    `gorm:"not null" json:"address"`
	Photo           string    `json:"photo"`
	Email           string    `gorm:"uniqueIndex;type:varchar(255);not null" json:"email"`
	EmailVerifiedAt time.Time `gorm:"<-:create"`
	Password        string    `gorm:"->;<-;not null" json:"-"`
	Role            string    `gorm:"type:ENUM('admin','member');not null" json:"role"`
	DeletedAt       gorm.DeletedAt
	CreatedAt       time.Time `gorm:"<-:create"`
	UpdatedAt       time.Time
	Token           string `json:"token,omitempty"`
}
