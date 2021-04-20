package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/nashirkra/RENTAL-BUKU/dto"
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"github.com/nashirkra/RENTAL-BUKU/repository"
)

type LoanService interface {
	Insert(lo dto.LoanCreate) entity.Loan
	Update(lo dto.LoanUpdate) entity.Loan
	Delete(lo entity.Loan)
	All() []entity.Loan
	FindByID(loanID uint64) entity.Loan
	UserRole(userID string) string
	CheckBookStock(bookID uint64) bool
	ReturnBook(lo dto.LoanUpdate) entity.Loan
	FinePayment(loanID uint64, xdays float64) bool
}

type loanService struct {
	loanRepository repository.LoanRepository
}

func NewLoanService(loanRepo repository.LoanRepository) LoanService {
	return &loanService{
		loanRepository: loanRepo,
	}
}

func (serv *loanService) Insert(lo dto.LoanCreate) entity.Loan {
	loan := entity.Loan{}
	err := smapping.FillStruct(&loan, smapping.MapFields(&lo))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := serv.loanRepository.InsertLoan(loan)
	return res
}
func (serv *loanService) Update(lo dto.LoanUpdate) entity.Loan {
	loan := entity.Loan{}
	err := smapping.FillStruct(&loan, smapping.MapFields(&lo))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := serv.loanRepository.UpdateLoan(loan)
	return res
}
func (serv *loanService) Delete(lo entity.Loan) {
	serv.loanRepository.DeleteLoan(lo)
}
func (serv *loanService) All() []entity.Loan {
	return serv.loanRepository.GetAllLoan()
}
func (serv *loanService) FindByID(loanID uint64) entity.Loan {
	return serv.loanRepository.FindLoanByID(loanID)
}
func (serv *loanService) UserRole(userID string) string {
	res := serv.loanRepository.UserRole(userID)
	return res
}
func (serv *loanService) CheckBookStock(bookID uint64) bool {
	return serv.loanRepository.CheckBookStock(bookID)
}
func (serv *loanService) FinePayment(loanID uint64, xdays float64) bool {
	return serv.loanRepository.FinePayment(loanID, xdays)
}
func (serv *loanService) ReturnBook(lo dto.LoanUpdate) entity.Loan {
	loan := entity.Loan{}
	err := smapping.FillStruct(&loan, smapping.MapFields(&lo))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := serv.loanRepository.ReturnBook(loan)
	return res
}
