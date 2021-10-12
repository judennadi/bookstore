package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/judennadi/bookstore/models"
	"github.com/judennadi/bookstore/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	nb := newBooks[0]
	fmt.Printf("%T", nb.ID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newBooks)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	if book.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Name must not be blank"})
		return
	} else if book.Author == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Author must not be blank"})
		return
	} else if book.Publication == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Publication must not be blank"})
		return
	}
	newBook, err := book.CreateBook()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errName := utils.HandleDuplicateError(err)

		message := fmt.Sprintf("%v already exist", errName)
		json.NewEncoder(w).Encode(map[string]string{"error": message})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	book, _ := models.GetBook(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	id := mux.Vars(r)["id"]
	oldBook, db := models.GetBook(id)
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
	// bookId, err := strconv.ParseInt(id, 0, 0)
	// if err != nil {
	// 	fmt.Println("error while parsing book id")
	// }
	book := models.DeleteBook(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
