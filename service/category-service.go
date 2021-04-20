package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/nashirkra/RENTAL-BUKU/dto"
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"github.com/nashirkra/RENTAL-BUKU/repository"
)

type CategoryService interface {
	Insert(c dto.CategoryCreate) entity.Category
	Update(c dto.CategoryUpdate) entity.Category
	Delete(c entity.Category)
	All() []entity.Category
	FindByID(catID uint64) entity.Category
	UserRole(userID string) string
}

type categoryService struct {
	catRepository repository.CategoryRepository
}

func NewCategoryService(catRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		catRepository: catRepo,
	}
}

func (serv *categoryService) Insert(c dto.CategoryCreate) entity.Category {
	cat := entity.Category{}
	err := smapping.FillStruct(&cat, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := serv.catRepository.InsertCategory(cat)
	return res
}
func (serv *categoryService) Update(c dto.CategoryUpdate) entity.Category {
	cat := entity.Category{}
	err := smapping.FillStruct(&cat, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := serv.catRepository.UpdateCategory(cat)
	return res
}
func (serv *categoryService) Delete(c entity.Category) {
	serv.catRepository.DeleteCategory(c)
}
func (serv *categoryService) All() []entity.Category {
	return serv.catRepository.GetAllCategory()
}
func (serv *categoryService) FindByID(catID uint64) entity.Category {
	return serv.catRepository.FindCategoryByID(catID)
}
func (serv *categoryService) UserRole(userID string) string {
	res := serv.catRepository.UserRole(userID)
	return res
}
