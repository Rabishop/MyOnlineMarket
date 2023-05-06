package model

type Tag_list struct {
	TagId   int64  `gorm:"column:tag_id;AUTO_INCREMENT;" json:"tagId"`
	TagName string `gorm:"column:tag_name;" json:"tagName"`
}

// Return TableName
func (Tag_list) TableName() string {
	return "tag_list"
}
