package model

type User struct {
	UserId       int64  `gorm:"column:user_id;AUTO_INCREMENT;" json:"userId"`
	UserAccount  string `gorm:"column:user_account;" json:"userAccount"`
	UserPassword string `gorm:"column:user_password;" json:"userPassword"`
	UserName     string `gorm:"column:user_name;" json:"userName"`
}

// Return TableName
func (User) TableName() string {
	return "user"
}
