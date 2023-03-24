package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// UserRegistResponse struct
type GameUploadResponse struct {
	Status string `json:"status"`
}

// UserRegistRequest struct
type GameUploadRequest struct {
	GamePrice    int64  `json:"gamePrice"`
	GameName     string `json:"gameName"`
	GameType     string `json:"gameType"`
	GameInfo     string `json:"gameInfo"`
	GameImg      string `json:"GameImg"`
	GameUploader string `json:"GameUploader"`
}

func UploadGame(gameUploadRequest *GameUploadRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var game model.Game
	game.GameUploader = gameUploadRequest.GameUploader
	game.GameImg = gameUploadRequest.GameImg
	game.GameInfo = gameUploadRequest.GameInfo
	game.GameName = gameUploadRequest.GameName
	game.GamePrice = gameUploadRequest.GamePrice
	game.GameType = gameUploadRequest.GameType

	if err := tx.Create(&game).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}

func UploadGameOutput(w http.ResponseWriter, gameUploadResponse *GameUploadResponse) {
	jsonbyte, err := json.Marshal(gameUploadResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
