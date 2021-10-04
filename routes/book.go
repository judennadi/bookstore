package routes

import (
	"github.com/gorilla/mux"
	"github.com/judennadi/bookstore/controllers"
)

// var router = mux.NewRouter().StrictSlash(true)
var BookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.EditBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
