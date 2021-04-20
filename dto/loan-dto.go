package dto

import (
	"time"

	"gorm.io/gorm"
)

type LoanCreate struct {
	UserID       uint64 `json:"user_id" form:"user_id" binding:"required"`
	BookID       uint64 `json:"book_id" form:"book_id" binding:"required"`
	BorrowedDate *time.Time
	DueDate      *time.Time
	ReturnDate   *time.Time
	DeletedAt    gorm.DeletedAt
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type LoanUpdate struct {
	ID           uint64 `json:"id" form:"id" binding:"required"`
	UserID       uint64 `json:"user_id" form:"user_id"`
	BookID       uint64 `json:"book_id" form:"book_id"`
	BorrowedDate *time.Time
	DueDate      *time.Time
	ReturnDate   *time.Time
	DeletedAt    gorm.DeletedAt
	UpdatedAt    time.Time
}
