package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// UserLoginResponse struct
type UserLogoutResponse struct {
	Status string `json:"status"`
}

// Delete the cookies to log out

func LogoutOutput(w http.ResponseWriter, userLoginResponse *UserLoginResponse) {
	jsonbyte, err := json.Marshal(userLoginResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
