package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/nashirkra/RENTAL-BUKU/dto"
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"github.com/nashirkra/RENTAL-BUKU/repository"
)

type BookService interface {
	Insert(b dto.BookCreate) entity.Book
	Update(b dto.BookUpdate) entity.Book
	Delete(b entity.Book)
	All() []entity.Book
	FindByID(bookID uint64) entity.Book
	IsAllowedToEdit(catID string, bookID uint64) bool
	UserRole(userID string) string
	FindBookByCat(catID uint64) []entity.Book
}

type bookService struct {
	bookRepository repository.BookRepository
}

//NewBookService .....
func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (service *bookService) Insert(b dto.BookCreate) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookRepository.InsertBook(book)
	return res
}

func (service *bookService) Update(b dto.BookUpdate) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookRepository.UpdateBook(book)
	return res
}

func (service *bookService) Delete(b entity.Book) {
	service.bookRepository.DeleteBook(b)
}

func (service *bookService) All() []entity.Book {
	return service.bookRepository.GetAllBook()
}

func (service *bookService) FindByID(bookID uint64) entity.Book {
	return service.bookRepository.FindBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(catID string, bookID uint64) bool {
	b := service.bookRepository.FindBookByID(bookID)
	id := fmt.Sprintf("%v", b.CategoryID)
	return catID == id
}

func (service *bookService) UserRole(userID string) string {
	res := service.bookRepository.UserRole(userID)
	return res
}

func (service *bookService) FindBookByCat(catID uint64) []entity.Book {
	return service.bookRepository.FindBookByCat(catID)
}
