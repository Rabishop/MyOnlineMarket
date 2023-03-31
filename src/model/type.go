package model

type Type struct {
	TypeId   int64  `gorm:"column:type_id;AUTO_INCREMENT;" json:"typeId"`
	TypeName string `gorm:"column:type_name;" json:"typeName"`
}

// Return TableName
func (Type) TableName() string {
	return "type"
}
