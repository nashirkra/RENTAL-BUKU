package entity

import (
	"time"

	"gorm.io/gorm"
)

/**
 * This is a Book entity class and is not intended for modification.
 */
type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:text" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Author      string `gorm:"type:text" json:"author"`
	Year        uint64 `gorm:"type:year" json:"year"`
	CategoryID  uint64 `gorm:"not null" json:"category_id"`
	Stock       uint64 `gorm:"type:int" json:"stock"`
	Status      uint   `gorm:"type:ENUM('1','0');not null" json:"status"`
	DeletedAt   gorm.DeletedAt
	CreatedAt   time.Time `gorm:"<-:create;not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	Category    Category  `gorm:"foreignkey:CategoryID;constraint:onUpdate:RESTRICT;onDelete:RESTRICT" json:"category"`
}
