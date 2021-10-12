package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/judennadi/bookstore/models"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		godotenv.Load()
		var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
		token, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid cookie"})
			return
		}
		parsedToken, err := jwt.Parse(token.Value, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid token"})
			return
		}
		claims := parsedToken.Claims.(jwt.MapClaims)
		email := fmt.Sprintf("%v", claims["email"])
		user := models.GetUserByEmail(email)
		fmt.Println(user)
		next.ServeHTTP(w, r)
	})
}

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
