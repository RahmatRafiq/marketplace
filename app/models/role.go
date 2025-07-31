package models

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Role string `json:"role"`
}
