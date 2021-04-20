package repository

import (
	"strconv"

	"github.com/nashirkra/RENTAL-BUKU/entity"
	"gorm.io/gorm"
)

type LoanRepository interface {
	InsertLoan(loan entity.Loan) entity.Loan
	UpdateLoan(loan entity.Loan) entity.Loan
	DeleteLoan(loan entity.Loan)
	GetAllLoan() []entity.Loan
	FindLoanByID(loanID uint64) entity.Loan
	UserRole(userID string) string
	CheckBookStock(bookID uint64) bool
	ReturnBook(loan entity.Loan) entity.Loan
	FinePayment(loanID uint64, xdays float64) bool
}

type loanConnection struct {
	connection *gorm.DB
}

/**
 * Create New LoanRepository
 */
func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanConnection{
		connection: db,
	}
}

func (db *loanConnection) InsertLoan(loan entity.Loan) entity.Loan {
	db.connection.Save(&loan)
	db.connection.Preload("User").Preload("Book").Find(&loan)
	return loan
}
func (db *loanConnection) UpdateLoan(loan entity.Loan) entity.Loan {
	db.connection.Save(&loan)
	db.connection.Preload("User").Preload("Book").Find(&loan)
	return loan
}
func (db *loanConnection) DeleteLoan(loan entity.Loan) {
	db.connection.Delete(&loan)
}
func (db *loanConnection) GetAllLoan() []entity.Loan {
	var loans []entity.Loan
	db.connection.Preload("User").Preload("Book").Find(&loans)
	return loans
}
func (db *loanConnection) FindLoanByID(loanID uint64) entity.Loan {
	var loan entity.Loan
	db.connection.Preload("User").Preload("Book").Find(&loan, loanID)
	return loan
}
func (db *loanConnection) UserRole(userID string) string {
	var user entity.User
	db.connection.Find(&user, userID)
	return user.Role
}
func (db *loanConnection) CheckBookStock(bookID uint64) bool {
	var book entity.Book
	db.connection.Find(&book, bookID)
	if book.Stock > 0 {
		book.Stock--
		db.connection.Save(&book)
		return true
	}
	return false
}
func (db *loanConnection) FinePayment(loanID uint64, xdays float64) bool {
	var fineP entity.FinePayment
	fineP.LoanID = loanID
	fineP.Amount = xdays * 5000
	fineP.Receipt = strconv.FormatInt(int64(xdays), 10) + " days delay"
	db.connection.Save(&fineP)
	return true
}

func (db *loanConnection) ReturnBook(loan entity.Loan) entity.Loan {
	var book entity.Book
	db.connection.Find(&book, loan.BookID)
	book.Stock++
	db.connection.Save(&book)
	db.connection.Save(&loan)
	db.connection.Preload("User").Preload("Book").Find(&loan)
	return loan
}
