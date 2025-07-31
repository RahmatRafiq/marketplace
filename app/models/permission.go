package models

type Permission struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Permission string `json:"permission"`
}
