package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// GameIndexResponse struct
type GameDetailsResponse struct {
	Status    string     `json:"status"`
	GameItem  model.Game `json:"gameItem"`
	Inventory bool       `json:"inventory"`
}

// GameIndexRequest struct
type GameDetailsRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
	GameName     string `json:"gameName"`
}

func GameDetails(gameDetailsRequest *GameDetailsRequest, gameDetailsResponse *GameDetailsResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	var userID model.UserID
	user.UserAccount = gameDetailsRequest.UserAccount
	user.UserPassword = gameDetailsRequest.UserPassword

	// get userID
	if err := tx.Model(&user).Where("user_account = ? AND user_password = ?", user.UserAccount, user.UserPassword).Take(&userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// get games by gameName
	if err := tx.Where("game_name = ?", gameDetailsRequest.GameName).Take(&gameDetailsResponse.GameItem).Error; err != nil {
		tx.Rollback()
		return err
	}

	var count int64
	if err := tx.Debug().Where("user_id = ? AND game_id = ?", userID.UserId, gameDetailsResponse.GameItem.GameId).Find(&model.Inventory{}).Count(&count).Error; err != nil {
		tx.Rollback()
		return err
	}

	// fmt.Println(count)

	if count > 0 {
		gameDetailsResponse.Inventory = true
	}

	tx.Commit()
	return tx.Error
}

func GameDetailsOutput(w http.ResponseWriter, gameDetailsResponse *GameDetailsResponse) {
	jsonbyte, err := json.Marshal(gameDetailsResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
