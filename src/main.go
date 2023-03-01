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

	var userRegistRequest user.UserRegistRequest
	var userRegistResponse user.UserRegistResponse
	userRegistResponse.Status = "Accepted"

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method != "POST" {
		userRegistResponse.Status = "Wrong Method"
		fmt.Println(userRegistResponse.Status)
		fmt.Println(r.Method)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	json.Unmarshal([]byte(body), &userRegistRequest)
	fmt.Println(r.Body)

	// Call Function
	err = user.Regist(&userRegistRequest)
	if err != nil {
		log.Println(err)
		return
	}

	// Return JSON
	jsonbyte, err := json.Marshal(userRegistResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}

// UserLogin API
func UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	var userLoginRequest user.UserLoginRequest
	var userLoginResponse user.UserLoginResponse
	userLoginResponse.Status = "Accepted"

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method != "POST" {
		userLoginResponse.Status = "Wrong Method"
		fmt.Println(userLoginResponse.Status)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	json.Unmarshal([]byte(body), &userLoginRequest)

	// Call Function
	err = user.Login(&userLoginRequest)
	if err != nil {
		log.Println(err)
		userLoginResponse.Status = "Password Error or Account Not Found"
	}

	// Return JSON
	jsonbyte, err := json.Marshal(userLoginResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}

func main() {

	http.HandleFunc("/user/regist", UserRegistHandler)

	http.HandleFunc("/user/login", UserLoginHandler)

	http.ListenAndServe(":8080", nil)
}
