package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/judennadi/bookstore/models"
	"github.com/judennadi/bookstore/utils"
)

// var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newBooks)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	newBook := book.CreateBook()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing book id")
	}
	book, _ := models.GetBook(bookId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	id := mux.Vars(r)["id"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing book id")
	}
	oldBook, db := models.GetBook(bookId)
	if newBook.Name != "" {
		oldBook.Name = newBook.Name
	}
	if newBook.Author != "" {
		oldBook.Author = newBook.Author
	}
	if newBook.Publication != "" {
		oldBook.Publication = newBook.Publication
	}
	db.Save(&oldBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(oldBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing book id")
	}
	book := models.DeleteBook(bookId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
