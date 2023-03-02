package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/m/v2/module/user"
)

// UserRegist API
func UserRegistHandler(w http.ResponseWriter, r *http.Request) {

	var userRegistRequest user.UserRegistRequest
	var userRegistResponse user.UserRegistResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userRegistResponse.Status = "Wrong Method"
		user.RegistOutput(w, &userRegistResponse)
		return
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &userRegistRequest)

	// Call Function
	err = user.Regist(&userRegistRequest)
	if err != nil {
		log.Println(err)
		userRegistResponse.Status = "Account already exists"
		user.RegistOutput(w, &userRegistResponse)
		return
	}

	// Return JSON
	userRegistResponse.Status = "Accepted"
	user.RegistOutput(w, &userRegistResponse)
}

// UserLogin API
func UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	var userLoginRequest user.UserLoginRequest
	var userLoginResponse user.UserLoginResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userLoginResponse.Status = "Wrong Method"
		user.LoginOutput(w, &userLoginResponse)
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
		userLoginResponse.Status = "Wrong Username or Password"

		// Set Cookie
		c1 := http.Cookie{
			Name:     "first_cookie",
			Value:    "vanyar",
			HttpOnly: true,
		}

		w.Header().Set("Set-Cookie", c1.String())

		user.LoginOutput(w, &userLoginResponse)
		return
	}

	// Call Function
	userLoginResponse.Status = "Accepted"
	user.LoginOutput(w, &userLoginResponse)

}

func main() {

	http.HandleFunc("/user/regist", UserRegistHandler)

	http.HandleFunc("/user/login", UserLoginHandler)

	fs := http.FileServer(http.Dir("../pages/"))
	http.Handle("/pages/", http.StripPrefix("/pages/", fs))

	http.ListenAndServe(":8080", nil)
}
