package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/model"
)

// CartCheckRequest struct
type CartCheckRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

// CartCheckResponse struct
type CartCheckResponse struct {
	Status string `json:"status"`
}

func CartCheck(cartCheckRequest *CartCheckRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	var userID model.UserID
	user.UserAccount = cartCheckRequest.UserAccount
	user.UserPassword = cartCheckRequest.UserPassword

	// get userID
	if err := tx.Model(&user).Where("user_account = ? AND user_password = ?", user.UserAccount, user.UserPassword).Take(&userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	cart := new([]model.Cart)

	// get the cart by userID
	if err := tx.Where("user_id = ?", userID.UserId).Find(cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Add games to a inventory
	inventory := new([]model.Inventory)
	for i := 0; i < len(*cart); i++ {
		var item model.Inventory
		item.GameId = (*cart)[i].GameId
		item.UserId = (*cart)[i].UserId
		item.InventoryDateAdded = time.Now().String()[0:20]

		*inventory = append(*inventory, item)
	}

	if err := tx.Create(inventory).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete games from the cart
	if err := tx.Where("user_id = ?", userID.UserId).Delete(&cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}

func CartCheckOutput(w http.ResponseWriter, cartCheckResponse *CartCheckResponse) {
	jsonbyte, err := json.Marshal(cartCheckResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
