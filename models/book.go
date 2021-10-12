package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/judennadi/bookstore/config"
)

var db *gorm.DB

type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

type Book struct {
	Model
	Name        string `json:"name" gorm:"not null;unique"`
	Author      string `json:"author" gorm:"not null"`
	Publication string `json:"publication" gorm:"not null"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &User{})
}

func (b *Book) CreateBook() (*Book, error) {
	result := db.Create(&b)
	return b, result.Error
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBook(id string) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id string) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
