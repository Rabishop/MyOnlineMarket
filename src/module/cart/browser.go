package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// CartBrowerResponse struct
type CartBrowserResponse struct {
	Status   string       `json:"status"`
	CartList []model.Cart `json:"cartList"`
	GameList []model.Game `json:"gameList"`
}

// CartBrowerRequest struct
type CartBrowserResqust struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

func CartBrowser(cartBrowerResqust *CartBrowserResqust, cartBrowerResponse *CartBrowserResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	var userID model.UserID
	gameItem := new(model.Game)
	user.UserAccount = cartBrowerResqust.UserAccount
	user.UserPassword = cartBrowerResqust.UserPassword

	// get userID
	if err := tx.Model(&user).Where("user_account = ? AND user_password = ?", user.UserAccount, user.UserPassword).Take(&userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	//　get gameID in the cart
	if err := tx.Where("user_id = ?", userID.UserId).Find(&cartBrowerResponse.CartList).Error; err != nil {
		tx.Rollback()
		return err
	}

	// get gameList by gameID
	for i := 0; i < len(cartBrowerResponse.CartList); i++ {
		if err := tx.Where("game_id = ?", cartBrowerResponse.CartList[i].GameId).Take(&gameItem).Error; err != nil {
			tx.Rollback()
			return err
		}
		// fmt.Println(gameItem)
		cartBrowerResponse.GameList = append(cartBrowerResponse.GameList, *gameItem)
	}

	// fmt.Println(cartBrowerResponse)

	tx.Commit()
	return tx.Error
}

func CartBrowserOutput(w http.ResponseWriter, cartBrowerResponse *CartBrowserResponse) {
	jsonbyte, err := json.Marshal(cartBrowerResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
