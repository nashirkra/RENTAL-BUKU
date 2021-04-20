package dto

import (
	"time"

	"gorm.io/gorm"
)

type CategoryUpdate struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
	DeletedAt gorm.DeletedAt
	UpdatedAt time.Time
}
type CategoryCreate struct {
	Name      string `json:"name" form:"name" binding:"required"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
