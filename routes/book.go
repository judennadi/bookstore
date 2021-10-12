package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/judennadi/bookstore/controllers"
	"github.com/judennadi/bookstore/middleware"
)

var BookStoreRoutes = func() {
	var router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.EditBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
	router.Use(middleware.Auth)
	http.Handle("/book/", router)
}
