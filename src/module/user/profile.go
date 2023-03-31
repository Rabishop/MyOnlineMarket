package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// UserLoginResponse struct
type UserProfileResponse struct {
	Status        string `json:"status"`
	UserName      string `json:"userName"`
	UserGameCount int    `json:"userGameCount"`
	UserPortrait  string `json:"userPortrait"`
}

// UserLoginResponse struct
type UserUploadPortraitRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
	UserPortrait string `json:"userPortrait"`
}

// UserLoginResponse struct
type UserUploadPortraitResponse struct {
	Status string `json:"status"`
}

func Profile(userLoginRequest *UserLoginRequest, userProfileResponse *UserProfileResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	user.UserAccount = userLoginRequest.UserAccount
	user.UserPassword = userLoginRequest.UserPassword

	if err := tx.Where("user_account = ? AND user_password = ?", user.UserAccount, user.UserPassword).Take(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	userProfileResponse.UserName = user.UserName
	userProfileResponse.UserPortrait = user.UserPortrait
	userProfileResponse.UserGameCount = 0

	tx.Commit()
	return tx.Error
}

func ProfileOutput(w http.ResponseWriter, userProflieResponse *UserProfileResponse) {
	jsonbyte, err := json.Marshal(userProflieResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}

func UploadPortrait(userUploadPortraitRequest *UserUploadPortraitRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	user.UserAccount = userUploadPortraitRequest.UserAccount
	user.UserPassword = userUploadPortraitRequest.UserPassword

	if err := tx.Where("user_account = ? AND user_password = ?", user.UserAccount, user.UserPassword).Model(&user).Update("user_portrait", userUploadPortraitRequest.UserPortrait).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}

func UploadPortraitOutput(w http.ResponseWriter, userUploadPortraitResponse *UserUploadPortraitResponse) {
	jsonbyte, err := json.Marshal(userUploadPortraitResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
