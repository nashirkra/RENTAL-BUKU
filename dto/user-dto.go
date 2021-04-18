package dto

import (
	"time"

	"gorm.io/gorm"
)

type UserUpdate struct {
	ID              uint64 `json:"id" form:"id" binding:"required"`
	Name            string `json:"name" form:"name" binding:"required"`
	Address         string `json:"address" form:"address" binding:"required"`
	Photo           string `json:"photo" form:"photo" binding:"required"`
	Email           string `json:"email" form:"email" binding:"required" validate:"email"`
	EmailVerifiedAt time.Time
	Password        string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
	Role            uint   `json:"role" form:"role" binding:"required"`
	DeletedAt       gorm.DeletedAt
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
