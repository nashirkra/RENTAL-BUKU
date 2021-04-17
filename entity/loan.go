package entity

import (
	"time"

	"gorm.io/gorm"
)

/**
 * This is a Loan entity class and is not intended for modification.
 */
type Loan struct {
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID       uint64 `gorm:"not null" json:"-"`
	BookID       uint64 `gorm:"not null" json:"-"`
	BorrowedDate *time.Time
	DueDate      *time.Time
	ReturnDate   *time.Time
	DeletedAt    gorm.DeletedAt
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         User `gorm:"foreignkey:UserID;constraint:onUpdate:RESTRICT;onDelete:RESTRICT" json:"user"`
	Book         Book `gorm:"foreignkey:BookID;constraint:onUpdate:RESTRICT;onDelete:RESTRICT" json:"book"`
}
