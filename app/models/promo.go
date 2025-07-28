package models

type Promo struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
