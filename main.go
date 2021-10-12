package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/judennadi/bookstore/controllers"
	"github.com/judennadi/bookstore/routes"
)

func main() {
	godotenv.Load()
	var router = mux.NewRouter().StrictSlash(true)
	routes.BookStoreRoutes()
	routes.FoodRoutes()
	routes.AuthRoutes(router)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Welcome</h1>")
	})
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {})
	http.Handle("/", router)
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(port, nil))
}
