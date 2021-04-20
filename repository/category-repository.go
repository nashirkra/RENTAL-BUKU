package repository

import (
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(cat entity.Category) entity.Category
	UpdateCategory(cat entity.Category) entity.Category
	DeleteCategory(cat entity.Category)
	GetAllCategory() []entity.Category
	FindCategoryByID(catID uint64) entity.Category
	UserRole(userID string) string
}

type categoryConnection struct {
	connection *gorm.DB
}

/**
 * Create New CategoryRepository
 */
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: db,
	}
}

func (db *categoryConnection) InsertCategory(cat entity.Category) entity.Category {
	db.connection.Save(&cat)
	return cat
}
func (db *categoryConnection) UpdateCategory(cat entity.Category) entity.Category {
	db.connection.Save(&cat)
	db.connection.Find(&cat)
	return cat
}
func (db *categoryConnection) DeleteCategory(cat entity.Category) {
	db.connection.Delete(&cat)
}
func (db *categoryConnection) GetAllCategory() []entity.Category {
	var cats []entity.Category
	db.connection.Find(&cats)
	return cats
}
func (db *categoryConnection) FindCategoryByID(catID uint64) entity.Category {
	var cat entity.Category
	db.connection.Find(&cat, catID)
	return cat
}
func (db *categoryConnection) UserRole(userID string) string {
	var user entity.User
	db.connection.Find(&user, userID)
	return user.Role
}
