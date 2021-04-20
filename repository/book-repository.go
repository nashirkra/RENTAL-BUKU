package repository

import (
	"github.com/nashirkra/RENTAL-BUKU/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(b entity.Book) entity.Book
	UpdateBook(b entity.Book) entity.Book
	DeleteBook(b entity.Book)
	GetAllBook() []entity.Book
	FindBookByID(bookID uint64) entity.Book
	UserRole(userID string) string
}

type bookConnection struct {
	connection *gorm.DB
}

/**
 * Create New BookRepository
 */
func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookConnection{
		connection: db,
	}
}

func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("Category").Find(&b)
	return b
}
func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)
	db.connection.Preload("Category").Find(&b)
	return b
}
func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b)
}
func (db *bookConnection) GetAllBook() []entity.Book {
	var books []entity.Book
	db.connection.Preload("Category").Find(&books)
	return books
}
func (db *bookConnection) FindBookByID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("Category").Find(&book, bookID)
	return book
}

func (db *bookConnection) UserRole(userID string) string {
	var user entity.User
	db.connection.Find(&user, userID)
	return user.Role
}
