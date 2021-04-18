package dto

import (
	"time"

	"gorm.io/gorm"
)

type Register struct {
	Name            string `json:"name" form:"name" binding:"required"`
	Address         string `json:"address" form:"address" binding:"required"`
	Photo           string `json:"photo" form:"photo" binding:"required"`
	Email           string `json:"email" form:"email" binding:"required,email" validate:"email"`
	EmailVerifiedAt time.Time
	Password        string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:""`
	Role            string `json:"role" form:"role"`
	DeletedAt       gorm.DeletedAt
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
