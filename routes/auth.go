package routes

import (
	"github.com/gorilla/mux"
	"github.com/judennadi/bookstore/controllers"
)

var AuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")
}
