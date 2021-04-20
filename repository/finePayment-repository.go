package repository

import (
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"gorm.io/gorm"
)

type FinePaymentRepository interface {
	InsertFinePayment(fp entity.FinePayment) entity.FinePayment
	UpdateFinePayment(fp entity.FinePayment) entity.FinePayment
	DeleteFinePayment(fp entity.FinePayment)
	GetAllFinePayment() []entity.FinePayment
	FindFinePaymentByID(fpID uint64) entity.FinePayment
	UserRole(userID string) string
}

type fpConnection struct {
	connection *gorm.DB
}

/**
 * Create New FinePaymentRepository
 */
func NewFinePaymentRepository(db *gorm.DB) FinePaymentRepository {
	return &fpConnection{
		connection: db,
	}
}

func (db *fpConnection) InsertFinePayment(fp entity.FinePayment) entity.FinePayment {
	db.connection.Save(&fp)
	db.connection.Preload("Loan").Save(&fp)
	return fp
}
func (db *fpConnection) UpdateFinePayment(fp entity.FinePayment) entity.FinePayment {
	db.connection.Save(&fp)
	db.connection.Preload("Loan").Save(&fp)
	return fp
}
func (db *fpConnection) DeleteFinePayment(fp entity.FinePayment) {
	db.connection.Delete(&fp)
}
func (db *fpConnection) GetAllFinePayment() []entity.FinePayment {
	var fps []entity.FinePayment
	db.connection.Preload("Loan").Find(&fps)
	return fps
}
func (db *fpConnection) FindFinePaymentByID(fpID uint64) entity.FinePayment {
	var fp entity.FinePayment
	db.connection.Preload("Loan").Find(&fp, fpID)
	return fp
}
func (db *fpConnection) UserRole(userID string) string {
	var user entity.User
	db.connection.Find(&user, userID)
	return user.Role
}
