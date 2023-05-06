package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/model"
)

// CartUploadRequest struct
type CartUploadRequest struct {
	UserAccount   string `json:"userAccount"`
	UserPassword  string `json:"userPassword"`
	GameId        int64  `json:"gameId"`
	CartDateAdded string `json:"cartDateAdded"`
}

// CartUploadResponse struct
type CartUploadResponse struct {
	Status string `json:"status"`
}

func CartUpload(cartUploadRequest *CartUploadRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	var userID model.UserID
	user.UserAccount = cartUploadRequest.UserAccount
	user.UserPassword = cartUploadRequest.UserPassword

	// get userID
	if err := tx.Model(&user).Where("user_account = ? AND user_password = ?", user.UserAccount, user.UserPassword).Take(&userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	cart := new(model.Cart)
	cart.UserId = userID.UserId
	cart.GameId = cartUploadRequest.GameId
	cart.CartDateAdded = time.Now().String()[0:20]

	// Upload games to a cart
	if err := tx.Create(&cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}

func CartUploadOutput(w http.ResponseWriter, cartUploadResponse *CartUploadResponse) {
	jsonbyte, err := json.Marshal(cartUploadResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
