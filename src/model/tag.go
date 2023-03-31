package model

type Tag struct {
	TagId    int64  `gorm:"column:tag_id;AUTO_INCREMENT;" json:"tagId"`
	GameId   int64  `gorm:"column:game_id;" json:"gameId"`
	GameName string `gorm:"column:game_name;" json:"gameName"`
	TagName  string `gorm:"column:tag_name;" json:"tagName"`
}

// Return TableName
func (Tag) TableName() string {
	return "tag"
}
