package user

import (
	"example.com/m/v2/model"
)

// UserLoginResponse struct
type UserLoginResponse struct {
	Status string `json:"status"`
}

// UserRegistRequest struct
type UserLoginRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

func Login(userLoginRequest *UserLoginRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	user.UserAccount = userLoginRequest.UserAccount
	user.UserPassword = userLoginRequest.UserPassword

	if err := tx.Debug().Where("user_account = ? AND user_password = ?", user.UserAccount, user.UserPassword).Take(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}
