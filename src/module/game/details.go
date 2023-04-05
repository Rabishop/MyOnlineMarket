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
	Status   string     `json:"status"`
	GameItem model.Game `json:"gameItem"`
}

// GameIndexRequest struct
type GameDetailsRequest struct {
	GameName string `json:"gameName"`
}

func GameDetails(gameDetailsRequest *GameDetailsRequest, gameDetailsResponse *GameDetailsResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	if err := tx.Debug().Where("game_name = ?", gameDetailsRequest.GameName).Take(&gameDetailsResponse.GameItem).Error; err != nil {
		tx.Rollback()
		return err
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
