package model

type Game struct {
	UserId       int64  `gorm:"column:user_id;AUTO_INCREMENT;" json:"user_id"`
	UserAccount  string `gorm:"column:user_account;" json:"user_account"`
	UserPassword string `gorm:"column:user_password;" json:"user_password"`
	UserName     string `gorm:"column:user_name;" json:"user_name"`
}

// Return TableName
func (Game) TableName() string {
	return "game"
}
