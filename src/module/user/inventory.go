package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// UserInventoryResponse struct
type UserInventoryResponse struct {
	Status   string            `json:"status"`
	GameList []model.Inventory `json:"gameList"`
	ItemList []model.Game      `json:"itemList"`
}

// UserInventoryRequest struct
type UserInventoryRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

func Inventory(userInventoryRequest *UserInventoryRequest, userInventoryResponse *UserInventoryResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	var userID model.UserID
	user.UserAccount = userInventoryRequest.UserAccount
	user.UserPassword = userInventoryRequest.UserPassword
	gameItem := new(model.Game)

	// get userID
	if err := tx.Model(&user).Where("user_account = ? AND user_password = ?", user.UserAccount, user.UserPassword).Take(&userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// get user's games
	if err := tx.Where("user_id = ?", userID.UserId).Find(&userInventoryResponse.GameList).Error; err != nil {
		tx.Rollback()
		return err
	}

	for i := 0; i < len(userInventoryResponse.GameList); i++ {
		if err := tx.Where("game_id = ?", userInventoryResponse.GameList[i].GameId).Take(&gameItem).Error; err != nil {
			tx.Rollback()
			return err
		}
		// fmt.Println(gameItem)
		userInventoryResponse.ItemList = append(userInventoryResponse.ItemList, *gameItem)
	}

	// fmt.Println(userInventoryResponse.GameList)

	tx.Commit()
	return tx.Error
}

func InventoryOutput(w http.ResponseWriter, userInventoryResponse *UserInventoryResponse) {
	jsonbyte, err := json.Marshal(userInventoryResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
