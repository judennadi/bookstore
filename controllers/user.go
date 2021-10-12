package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/judennadi/bookstore/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetUsers()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUsers)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userId := uuid.MustParse(id)
	// userId, err := strconv.ParseInt(id, 0, 0)
	// if err != nil {
	// 	fmt.Println("error while parsing user id")
	// }
	user := models.DeleteUser(userId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"user deleted": user})
}
