package model

type Tag struct {
	Tag_id   int64  `gorm:"column:tag_id;AUTO_INCREMENT;" json:"tagId"`
	Game_id  int64  `gorm:"column:game_id;" json:"gameId"`
	Tag_Name string `gorm:"column:tag_name;" json:"tagName"`
}

// Return TableName
func (Tag) TableName() string {
	return "tag"
}
