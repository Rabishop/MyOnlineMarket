package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// CartRemoveRequest struct
type CartRemoveRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
	GameId       int64  `json:"gameId"`
}

// CartRemoveResponse struct
type CartRemoveResponse struct {
	Status string `json:"status"`
}

func CartRemove(cartRemoveRequest *CartRemoveRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	var userID model.UserID
	user.UserAccount = cartRemoveRequest.UserAccount
	user.UserPassword = cartRemoveRequest.UserPassword

	// get userID
	if err := tx.Model(&user).Where("user_account = ? AND user_password = ?", user.UserAccount, user.UserPassword).Take(&userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	cart := new(model.Cart)
	cart.UserId = userID.UserId
	cart.GameId = cartRemoveRequest.GameId

	// Delete games
	if err := tx.Where("user_id = ? AND game_id = ?", cart.UserId, cart.GameId).Delete(&cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}

func CartRemoveOutput(w http.ResponseWriter, cartRemoveResponse *CartRemoveResponse) {
	jsonbyte, err := json.Marshal(cartRemoveResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
