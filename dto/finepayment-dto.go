package dto

import (
	"time"

	"gorm.io/gorm"
)

type FinePaymentCreate struct {
	Receipt   string  `json:"receipt" form:"receipt"`
	Amount    float64 `json:"amount" form:"amount"`
	LoanID    uint64  `json:"loan_id" form:"loan_id"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
type FinePaymentUpdate struct {
	ID        uint64  `json:"id" form:"id"`
	Receipt   string  `json:"receipt" form:"receipt"`
	Amount    float64 `json:"amount" form:"amount"`
	LoanID    uint64  `json:"loan_id" form:"loan_id"`
	DeletedAt gorm.DeletedAt
	UpdatedAt time.Time
}
