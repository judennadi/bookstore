package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
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
