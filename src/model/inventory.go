package model

type Inventory struct {
	UserId               int64  `gorm:"column:user_id;" json:"userId"`
	GameId               int64  `gorm:"column:game_id;" json:"gameId"`
	InventoryDateAdded   string `gorm:"column:inventory_date_added;" json:"InventoryDateAdded"`
	InventoryHoursPlayed int64  `gorm:"column:inventory_hours_played;" json:"InventoryHoursPlayed"`
}

// Return TableName
func (Inventory) TableName() string {
	return "inventory"
}
