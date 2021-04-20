package entity

import (
	"time"

	"gorm.io/gorm"
)

/**
 * This is a FinePayment entity class and is not intended for modification.
 */
type FinePayment struct {
	ID        uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Receipt   string  `gorm:"type:varchar(255)" json:"receipt"`
	Amount    float64 `gorm:"type:double" json:"amount"`
	LoanID    uint64  `gorm:"not null" json:"-"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time `gorm:"<-:create;not null"`
	UpdatedAt time.Time `gorm:"not null"`
	Loan      Loan      `gorm:"foreignkey:LoanID;constraint:onUpdate:RESTRICT;onDelete:RESTRICT" json:"loan"`
}
