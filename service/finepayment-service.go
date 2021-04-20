package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/nashirkra/RENTAL-BUKU/dto"
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"github.com/nashirkra/RENTAL-BUKU/repository"
)

type FinePaymentService interface {
	Insert(fp dto.FinePaymentCreate) entity.FinePayment
	Update(fp dto.FinePaymentUpdate) entity.FinePayment
	Delete(fp entity.FinePayment)
	All() []entity.FinePayment
	FindByID(fpID uint64) entity.FinePayment
	UserRole(userID string) string
}

type finePaymentService struct {
	finePaymentRepository repository.FinePaymentRepository
}

func NewFinePaymentService(fpRepo repository.FinePaymentRepository) FinePaymentService {
	return &finePaymentService{
		finePaymentRepository: fpRepo,
	}
}

func (serv *finePaymentService) Insert(fp dto.FinePaymentCreate) entity.FinePayment {
	fineP := entity.FinePayment{}
	err := smapping.FillStruct(&fineP, smapping.MapFields(&fp))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := serv.finePaymentRepository.InsertFinePayment(fineP)
	return res
}
func (serv *finePaymentService) Update(fp dto.FinePaymentUpdate) entity.FinePayment {
	fineP := entity.FinePayment{}
	err := smapping.FillStruct(&fineP, smapping.MapFields(&fp))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := serv.finePaymentRepository.UpdateFinePayment(fineP)
	return res
}
func (serv *finePaymentService) Delete(fp entity.FinePayment) {
	serv.finePaymentRepository.DeleteFinePayment(fp)
}
func (serv *finePaymentService) All() []entity.FinePayment {
	return serv.finePaymentRepository.GetAllFinePayment()
}
func (serv *finePaymentService) FindByID(fpID uint64) entity.FinePayment {
	return serv.finePaymentRepository.FindFinePaymentByID(fpID)
}
func (serv *finePaymentService) UserRole(userID string) string {
	res := serv.finePaymentRepository.UserRole(userID)
	return res
}
