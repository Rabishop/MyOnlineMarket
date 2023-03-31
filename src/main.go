package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/m/v2/module/game"
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
		userRegistResponse.Status = "SQL Access Error"
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
		userLoginResponse.Status = "SQL Access Error"
		user.LoginOutput(w, &userLoginResponse)
		return
	}

	// Set Cookie
	cookie1 := http.Cookie{Name: "userAccount", Value: userLoginRequest.UserAccount, Path: "/", MaxAge: 86400}
	cookie2 := http.Cookie{Name: "userPassword", Value: userLoginRequest.UserPassword, Path: "/", MaxAge: 86400}
	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)

	// Return JSON
	userLoginResponse.Status = "Accepted"
	user.LoginOutput(w, &userLoginResponse)
}

// UserLogout API
func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {

	var userLogoutResponse user.UserLoginResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userLogoutResponse.Status = "Wrong Method"
		user.LogoutOutput(w, &userLogoutResponse)
	}

	// Set Cookie
	cookie1 := http.Cookie{Name: "userAccount", Value: "", Path: "/", MaxAge: -1}
	cookie2 := http.Cookie{Name: "userPassword", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)

	// Return JSON
	userLogoutResponse.Status = "Accepted"
	user.LogoutOutput(w, &userLogoutResponse)
}

// UserProfile API
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {

	var userProflieResponse user.UserProfileResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userProflieResponse.Status = "Wrong Method"
		user.ProfileOutput(w, &userProflieResponse)
	}
	// Read Cookie
	cookie1, err1 := r.Cookie("userAccount")
	cookie2, err2 := r.Cookie("userPassword")
	if err1 != nil || err2 != nil {
		log.Println(err1)
		log.Println(err2)
		userProflieResponse.Status = "SQL Access Error"
		user.ProfileOutput(w, &userProflieResponse)
		return
	}
	var userLoginRequest = user.UserLoginRequest{UserAccount: cookie1.Value, UserPassword: cookie2.Value}

	// Call Function
	err := user.Profile(&userLoginRequest, &userProflieResponse)
	if err != nil {
		log.Println(err)
		userProflieResponse.Status = "SQL Access Error"
		user.ProfileOutput(w, &userProflieResponse)
		return
	}

	// Return JSON
	userProflieResponse.Status = "Accepted"
	user.ProfileOutput(w, &userProflieResponse)
}

func UserUploadPortraitHandler(w http.ResponseWriter, r *http.Request) {

	var userUploadPortraitRequest user.UserUploadPortraitRequest
	var userUploadPortraitResponse user.UserUploadPortraitResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userUploadPortraitResponse.Status = "Wrong Method"
		user.UploadPortraitOutput(w, &userUploadPortraitResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &userUploadPortraitRequest)

	// Read Cookie
	cookie1, err1 := r.Cookie("userAccount")
	cookie2, err2 := r.Cookie("userPassword")
	if err1 != nil || err2 != nil {
		log.Println(err1)
		log.Println(err2)
		userUploadPortraitResponse.Status = "SQL Access Error"
		user.UploadPortraitOutput(w, &userUploadPortraitResponse)
		return
	}

	userUploadPortraitRequest.UserAccount = cookie1.Value
	userUploadPortraitRequest.UserPassword = cookie2.Value

	// fmt.Println(userUploadPortraitRequest)
	// Call Function
	err = user.UploadPortrait(&userUploadPortraitRequest)
	if err != nil {
		log.Println(err)
		userUploadPortraitResponse.Status = "SQL Access Error"
		user.UploadPortraitOutput(w, &userUploadPortraitResponse)
		return
	}

	// Return JSON
	userUploadPortraitResponse.Status = "Accepted"
	user.UploadPortraitOutput(w, &userUploadPortraitResponse)
}

func GameUploadHandler(w http.ResponseWriter, r *http.Request) {

	var gameUploadRequest game.GameUploadRequest
	var gameUploadResponse game.GameUploadResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		gameUploadResponse.Status = "Wrong Method"
		game.GameUploadOutput(w, &gameUploadResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &gameUploadRequest)

	// Read Cookie
	cookie1, err1 := r.Cookie("userAccount")
	if err1 != nil {
		log.Println(err1)
		gameUploadResponse.Status = "SQL Access Error"
		game.GameUploadOutput(w, &gameUploadResponse)
		return
	}

	gameUploadRequest.GameUploader = cookie1.Value

	// Call Function
	err = game.GameUpload(&gameUploadRequest)
	if err != nil {
		log.Println(err)
		gameUploadResponse.Status = "SQL Access Error"
		game.GameUploadOutput(w, &gameUploadResponse)
		return
	}

	// Return JSON
	gameUploadResponse.Status = "Accepted"
	game.GameUploadOutput(w, &gameUploadResponse)
}

func GameIndexHandler(w http.ResponseWriter, r *http.Request) {

	var gameIndexResponse game.GameIndexResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		gameIndexResponse.Status = "Wrong Method"
		game.GameIndexOutput(w, &gameIndexResponse)
	}

	// Call Function
	err := game.GameIndex(&gameIndexResponse)
	if err != nil {
		log.Println(err)
		gameIndexResponse.Status = "SQL Access Error"
		game.GameIndexOutput(w, &gameIndexResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	gameIndexResponse.Status = "Accepted"
	game.GameIndexOutput(w, &gameIndexResponse)
}

func main() {

	// Functions Handle

	http.HandleFunc("/user/regist", UserRegistHandler)

	http.HandleFunc("/user/login", UserLoginHandler)

	http.HandleFunc("/user/logout", UserLogoutHandler)

	http.HandleFunc("/user/profile", UserProfileHandler)

	http.HandleFunc("/user/uploadPortrait", UserUploadPortraitHandler)

	http.HandleFunc("/game/upload", GameUploadHandler)

	http.HandleFunc("/game/index", GameIndexHandler)

	// Pages Handle

	fs := http.FileServer(http.Dir("../"))

	http.Handle("/", http.StripPrefix("", fs))

	// Build the Server

	http.ListenAndServe(":8080", nil)
}
