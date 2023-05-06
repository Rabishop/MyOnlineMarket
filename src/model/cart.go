package model

type Cart struct {
	UserId        int64  `gorm:"column:user_id;" json:"userId"`
	GameId        int64  `gorm:"column:game_id;" json:"gameId"`
	CartDateAdded string `gorm:"column:cart_date_added;" json:"cartDateAdded"`
}

// Return TableName
func (Cart) TableName() string {
	return "cart"
}
