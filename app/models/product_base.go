package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductBase struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Products  []Product      `gorm:"foreignKey:ProductBaseID" json:"products"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}
