package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/judennadi/bookstore/controllers"
	"github.com/judennadi/bookstore/middleware"
)

var FoodRoutes = func() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user/", controllers.GetUsers).Methods("GET")
	// router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	// router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	// router.HandleFunc("/book/{id}", controllers.EditBook).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.Use(middleware.CheckAuth)
	http.Handle("/user/", router)
}
