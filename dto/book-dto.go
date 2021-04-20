package dto

import (
	"time"

	"gorm.io/gorm"
)

type BookUpdate struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Author      string `json:"author" form:"author"`
	Year        uint64 `json:"year" form:"year"`
	CategoryID  uint64 `json:"category_id" form:"category_id"`
	Stock       uint64 `json:"stock" form:"stock"`
	Status      uint   `json:"status" form:"status"`
	DeletedAt   gorm.DeletedAt
	UpdatedAt   time.Time
}

type BookCreate struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Author      string `json:"author" form:"author"`
	Year        uint64 `json:"year" form:"year"`
	CategoryID  uint64 `json:"category_id" form:"category_id"`
	Stock       uint64 `json:"stock" form:"stock"`
	Status      uint   `json:"status" form:"status"`
	DeletedAt   gorm.DeletedAt
	CreatedAt   time.Time
	UpdatedAt   time.Time
	//	Category    Category `gorm:"foreignkey:CategoryID;constraint:onUpdate:RESTRICT;onDelete:RESTRICT" json:"category"`
}

/*
type Category struct {
	ID        uint64
	Name      string
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
*/
