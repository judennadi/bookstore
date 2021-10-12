package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/judennadi/bookstore/models"
	"github.com/judennadi/bookstore/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var reqBody = &models.User{}
	json.NewDecoder(r.Body).Decode(reqBody)

	if reqBody.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Name must not be blank"})
		return
	} else if reqBody.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Email must not be blank"})
		return
	} else if reqBody.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Password must not be blank"})
		return
	}

	if err := reqBody.HashPassword(); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Couldn't hash password"})
		return
	}

	newUser, err := reqBody.CreateUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errName := utils.HandleDuplicateError(err)

		message := fmt.Sprintf("%v already exist", errName)
		json.NewEncoder(w).Encode(map[string]string{"error": message})
		return
	}

	token, jwtErr := utils.GenerateJWT(newUser.Email)
	if jwtErr != nil {
		json.NewEncoder(w).Encode(jwtErr)
		return
	}
	cookie := &http.Cookie{Name: "token", Value: token, HttpOnly: true, Expires: time.Now().Add(time.Minute * 5)}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(map[string]interface{}{"user": newUser})

}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqBody = &models.User{}
	json.NewDecoder(r.Body).Decode(reqBody)

	if reqBody.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Email must not be blank"})
		return
	} else if reqBody.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Password must not be blank"})
		return
	}

	newUser := models.GetUserByEmail(reqBody.Email)
	if newUser.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect username or password"})
		return
	}

	if err := newUser.ComparePassword(reqBody.Password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect username or password"})
		return
	}

	token, jwtErr := utils.GenerateJWT(newUser.Email)
	if jwtErr != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Couldn't generate token"})
		return
	}
	cookie := &http.Cookie{Name: "token", Value: token, HttpOnly: true, Expires: time.Now().Add(time.Hour * 12)}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(map[string]interface{}{"user": newUser})

}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{Name: "token", Value: "", HttpOnly: true, Expires: time.Now().Add(time.Minute - 1)}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(map[string]interface{}{"user": "logged out"})

}
