package entity

type Category struct {
	ID      uint   `gorm:"primary_key";"AUTO_INCREMENT" json:"id"`
	Name    string `gorm:"not null" json:"name"`
	Removed bool   `gorm:"default:false" json:"removed"`
}

func (c Category) TableName() string {
	return "category"
}