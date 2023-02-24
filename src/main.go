package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/m/v2/module/user"
)

// UserRegist API
func UserRegistHandler(w http.ResponseWriter, r *http.Request) {

	// Check Method
	if r.Method != "POST" {
		jsonbyte, err := json.Marshal(user.UserRegistResponse{"Wrong Method"})
		if err != nil {
			fmt.Println("Marshal failed")
		}
		fmt.Fprintln(w, string(jsonbyte))
	}

	// Read Body
	var userRegistRequest user.UserRegistRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &userRegistRequest)

	// Call Function
	user.Regist(&userRegistRequest)

	// jsonbyte, err := json.Marshal(UserRegistResponse{"Accepted"})
	// if err != nil {
	// 	fmt.Println("Marshal failed")
	// }

	// Return JSON
	fmt.Fprintln(w, string(jsonbyte))
}

func main() {

	// ユーザ所持キャラクター一覧取得API
	http.HandleFunc("/character/list", UserRegistHandler)

	http.ListenAndServe(":8080", nil)
}
