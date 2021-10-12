package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func ParseBody(r *http.Request, x interface{}) {
	// if body, err := ioutil.ReadAll(r.Body); err == nil {
	// 	if err := json.Unmarshal([]byte(body), x); err != nil {
	// 		return
	// 	}
	// }
	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
		return
	}
}

func HandleDuplicateError(e error) string {
	var newErrJson interface{}
	errJson, _ := json.Marshal(e)
	err := json.Unmarshal(errJson, &newErrJson)
	if err != nil {
		fmt.Println(err)
	}
	errMap := newErrJson.(map[string]interface{})
	code := fmt.Sprintf("%v", errMap["Code"])
	errCode, _ := strconv.Atoi(code)
	var m string
	if errCode == 23505 {
		a := fmt.Sprintf("%v", errMap["Detail"])
		b := strings.Index(a, ")")
		c := a[:b]
		d := strings.Index(c, "(")
		m = c[d+1:]
	}
	return m
}

func GenerateJWT(email string) (string, error) {
	godotenv.Load()
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	newJwt := jwt.New(jwt.SigningMethodHS256)
	claims := newJwt.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	token, err := newJwt.SignedString(jwtSecret)
	if err != nil {
		fmt.Printf("couldn't sign JWT: %v", err.Error())
		return "", err
	}
	return token, nil
}
