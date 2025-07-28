package models

type Warehouse struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
