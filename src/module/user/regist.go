package user

import (
	"example.com/m/v2/model"
)

// UserGetResponse struct
type UserRegistResponse struct {
	Status string `json:"status"`
}

// UserCreateRequest struct
type UserRegistRequest struct {
	UserAccount  string `json:"user_account"`
	UserPassword string `json:"user_password"`
	UserName     string `json:"user_name"`
}

func Regist(userRegistRequest *UserRegistRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	user := new(model.User)
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
