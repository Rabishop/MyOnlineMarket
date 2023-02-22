package user

import (
	"example.com/m/v2/model"
)

func Regist() error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	user := new(model.User)
	user.UserName = "Alice"
	user.UserAccount = "Alice123"
	user.UserPassword = "123456"

	if err := tx.Debug().Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}
