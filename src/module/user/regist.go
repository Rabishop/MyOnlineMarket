package user

import (
	"example.com/m/v2/model"
)

// UserRegistResponse struct
type UserRegistResponse struct {
	Status string `json:"status"`
}

// UserRegistRequest struct
type UserRegistRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
	UserName     string `json:"userName"`
}

func Regist(userRegistRequest *UserRegistRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	user.UserName = userRegistRequest.UserName
	user.UserAccount = userRegistRequest.UserAccount
	user.UserPassword = userRegistRequest.UserPassword

	if err := tx.Debug().Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}
