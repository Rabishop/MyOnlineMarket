package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/module/cart"
	"example.com/m/v2/module/game"
	"example.com/m/v2/module/user"
)

// UserRegist
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

// UserLogin
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

// UserLogout
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

// View user's Profile
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

// View user's inventory
func UserInventoryHandler(w http.ResponseWriter, r *http.Request) {

	var userInventoryRequest user.UserInventoryRequest
	var userInventoryResponse user.UserInventoryResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userInventoryResponse.Status = "Wrong Method"
		user.InventoryOutput(w, &userInventoryResponse)
	}
	// Read Cookie
	cookie1, err1 := r.Cookie("userAccount")
	cookie2, err2 := r.Cookie("userPassword")
	if err1 != nil || err2 != nil {
		log.Println(err1)
		log.Println(err2)
		userInventoryResponse.Status = "SQL Access Error"
		user.InventoryOutput(w, &userInventoryResponse)
		return
	}

	userInventoryRequest.UserAccount = cookie1.Value
	userInventoryRequest.UserPassword = cookie2.Value

	// Call Function
	err := user.Inventory(&userInventoryRequest, &userInventoryResponse)
	if err != nil {
		log.Println(err)
		userInventoryResponse.Status = "SQL Access Error"
		user.InventoryOutput(w, &userInventoryResponse)
		return
	}

	// Return JSON
	userInventoryResponse.Status = "Accepted"
	user.InventoryOutput(w, &userInventoryResponse)
}

// Upload user's portrait
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

// Upload games
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

// View games on the homepage
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

// View games by types
func GameBrowserHandler(w http.ResponseWriter, r *http.Request) {

	var gameBrowserResqust game.GameBrowserResqust
	var gameBrowserResponse game.GameBrowserResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		gameBrowserResponse.Status = "Wrong Method"
		game.GameBrowserOutput(w, &gameBrowserResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &gameBrowserResqust)

	// Call Function
	err = game.GameBrowser(&gameBrowserResqust, &gameBrowserResponse)
	if err != nil {
		log.Println(err)
		gameBrowserResponse.Status = "SQL Access Error"
		game.GameBrowserOutput(w, &gameBrowserResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	gameBrowserResponse.Status = "Accepted"
	game.GameBrowserOutput(w, &gameBrowserResponse)
}

// View game's details
func GameDetailsHandler(w http.ResponseWriter, r *http.Request) {

	var gameDetailsRequest game.GameDetailsRequest
	var gameDetailsResponse game.GameDetailsResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		gameDetailsResponse.Status = "Wrong Method"
		game.GameDetailsOutput(w, &gameDetailsResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &gameDetailsRequest)

	// Read Cookie
	cookie1, err1 := r.Cookie("userAccount")
	cookie2, err2 := r.Cookie("userPassword")
	if err1 != nil || err2 != nil {
		log.Println(err1)
		log.Println(err2)
		return
	}
	gameDetailsRequest.UserAccount = cookie1.Value
	gameDetailsRequest.UserPassword = cookie2.Value

	// Call Function
	err = game.GameDetails(&gameDetailsRequest, &gameDetailsResponse)
	if err != nil {
		log.Println(err)
		gameDetailsResponse.Status = "SQL Access Error"
		game.GameDetailsOutput(w, &gameDetailsResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	gameDetailsResponse.Status = "Accepted"
	game.GameDetailsOutput(w, &gameDetailsResponse)
}

// Add games to a cart
func CartUploadHandler(w http.ResponseWriter, r *http.Request) {

	var cartUploadRequest cart.CartUploadRequest
	var cartUploadResponse cart.CartUploadResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		cartUploadResponse.Status = "Wrong Method"
		cart.CartUploadOutput(w, &cartUploadResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &cartUploadRequest)

	// Read Cookie
	cookie1, err1 := r.Cookie("userAccount")
	cookie2, err2 := r.Cookie("userPassword")
	if err1 != nil || err2 != nil {
		log.Println(err1)
		log.Println(err2)
		return
	}
	cartUploadRequest.UserAccount = cookie1.Value
	cartUploadRequest.UserPassword = cookie2.Value
	cartUploadRequest.CartDateAdded = time.Now().String()

	fmt.Println(cartUploadRequest)

	// Call Function
	err = cart.CartUpload(&cartUploadRequest)
	if err != nil {
		log.Println(err)
		cartUploadResponse.Status = "SQL Access Error"
		cart.CartUploadOutput(w, &cartUploadResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	cartUploadResponse.Status = "Accepted"
	cart.CartUploadOutput(w, &cartUploadResponse)
}

// View the cart
func CartBrowserHandler(w http.ResponseWriter, r *http.Request) {

	var cartBrowserResqust cart.CartBrowserResqust
	var cartBrowserResponse cart.CartBrowserResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		cartBrowserResponse.Status = "Wrong Method"
		cart.CartBrowserOutput(w, &cartBrowserResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &cartBrowserResqust)

	// Read Cookie
	cookie1, err1 := r.Cookie("userAccount")
	cookie2, err2 := r.Cookie("userPassword")
	if err1 != nil || err2 != nil {
		log.Println(err1)
		log.Println(err2)
		return
	}
	cartBrowserResqust.UserAccount = cookie1.Value
	cartBrowserResqust.UserPassword = cookie2.Value

	// fmt.Println(cartBrowserResqust)

	// Call Function
	err = cart.CartBrowser(&cartBrowserResqust, &cartBrowserResponse)
	if err != nil {
		log.Println(err)
		cartBrowserResponse.Status = "SQL Access Error"
		cart.CartBrowserOutput(w, &cartBrowserResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	cartBrowserResponse.Status = "Accepted"
	cart.CartBrowserOutput(w, &cartBrowserResponse)
}

// Remove games from a cart
func CartRemoveHandler(w http.ResponseWriter, r *http.Request) {

	var cartRemoveRequest cart.CartRemoveRequest
	var cartRemoveResponse cart.CartRemoveResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		cartRemoveResponse.Status = "Wrong Method"
		cart.CartRemoveOutput(w, &cartRemoveResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &cartRemoveRequest)

	// Read Cookie
	cookie1, err1 := r.Cookie("userAccount")
	cookie2, err2 := r.Cookie("userPassword")
	if err1 != nil || err2 != nil {
		log.Println(err1)
		log.Println(err2)
		return
	}
	cartRemoveRequest.UserAccount = cookie1.Value
	cartRemoveRequest.UserPassword = cookie2.Value

	// Call Function
	err = cart.CartRemove(&cartRemoveRequest)
	if err != nil {
		log.Println(err)
		cartRemoveResponse.Status = "SQL Access Error"
		cart.CartRemoveOutput(w, &cartRemoveResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	cartRemoveResponse.Status = "Accepted"
	cart.CartRemoveOutput(w, &cartRemoveResponse)
}

// Add games to a inventory and Remove games from the cart
func CartCheckHandler(w http.ResponseWriter, r *http.Request) {

	var cartCheckRequest cart.CartCheckRequest
	var cartCheckResponse cart.CartCheckResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		cartCheckResponse.Status = "Wrong Method"
		cart.CartCheckOutput(w, &cartCheckResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &cartCheckRequest)

	// Read Cookie
	cookie1, err1 := r.Cookie("userAccount")
	cookie2, err2 := r.Cookie("userPassword")
	if err1 != nil || err2 != nil {
		log.Println(err1)
		log.Println(err2)
		return
	}
	cartCheckRequest.UserAccount = cookie1.Value
	cartCheckRequest.UserPassword = cookie2.Value

	// fmt.Println(cartRemoveRequest)

	// Call Function
	err = cart.CartCheck(&cartCheckRequest)
	if err != nil {
		log.Println(err)
		cartCheckResponse.Status = "SQL Access Error"
		cart.CartCheckOutput(w, &cartCheckResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	cartCheckResponse.Status = "Accepted"
	cart.CartCheckOutput(w, &cartCheckResponse)
}

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	// 获取请求的路径
// 	path := r.URL.Path

// 	// 检查路径中是否包含文件后缀
// 	if strings.HasSuffix(path, ".html") {
// 		// 如果有后缀，则重定向到去掉后缀的路径
// 		newPath := strings.TrimSuffix(path, ".html")
// 		http.Redirect(w, r, newPath, http.StatusMovedPermanently)
// 		return
// 	}
// }

func main() {

	// Functions Handle
	// http.HandleFunc("/", handleRequest)

	http.HandleFunc("/user/regist", UserRegistHandler)

	http.HandleFunc("/user/login", UserLoginHandler)

	http.HandleFunc("/user/logout", UserLogoutHandler)

	http.HandleFunc("/user/profile", UserProfileHandler)

	http.HandleFunc("/user/inventory", UserInventoryHandler)

	http.HandleFunc("/user/uploadPortrait", UserUploadPortraitHandler)

	http.HandleFunc("/game/upload", GameUploadHandler)

	http.HandleFunc("/game/index", GameIndexHandler)

	http.HandleFunc("/game/browser", GameBrowserHandler)

	http.HandleFunc("/game/details", GameDetailsHandler)

	http.HandleFunc("/cart/upload", CartUploadHandler)

	http.HandleFunc("/cart/browser", CartBrowserHandler)

	http.HandleFunc("/cart/remove", CartRemoveHandler)

	http.HandleFunc("/cart/check", CartCheckHandler)

	// Pages Handle

	fs := http.FileServer(http.Dir("../pages/"))

	http.Handle("/", http.StripPrefix("", fs))

	// Build the Server

	http.ListenAndServe(":8080", nil)
}
