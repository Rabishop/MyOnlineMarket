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

func GameUpload(gameUploadRequest *GameUploadRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	game := new(model.Game)
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

	tag := new(model.Tag)

	if err := tx.First(&game, "game_name = ?", game.GameName).Error; err != nil {
		tx.Rollback()
		return err
	}
	tag.GameId = game.GameId
	tag.GameName = game.GameName

	for i := 0; i < len(game.GameType); i++ {
		if game.GameType[i] == ';' {
			// fmt.Println(tag.TagName)

			ID := new(model.Type)
			if err := tx.Take(&ID, "type_name = ?", tag.TagName).Error; err != nil {
				tx.Rollback()
				return err
			}

			tag.TagId = ID.TypeId

			if err := tx.Create(&tag).Error; err != nil {
				tx.Rollback()
				return err
			}

			tag.TagName = ""
		} else {
			tag.TagName += string(game.GameType[i])
		}
	}

	tx.Commit()
	return tx.Error
}

func GameUploadOutput(w http.ResponseWriter, gameUploadResponse *GameUploadResponse) {
	jsonbyte, err := json.Marshal(gameUploadResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
