package seeds

import (
	"golang_starter_kit_2025/app/features/category"
	"log"
	"time"

	"gorm.io/gorm"
)

func SeedCategorySeeder(db *gorm.DB) error {
	log.Println("🌱 Seeding CategorySeeder...")

	categories := []category.Category{
		{Category: "Elektronik", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Category: "Fashion", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Category: "Makanan", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.Create(&categories).Error; err != nil {
		return err
	}
	return nil
}

func RollbackCategorySeeder(db *gorm.DB) error {
	log.Println("🗑️ Rolling back CategorySeeder…")
	return db.Unscoped().Where("category IN ?", []string{"Elektronik", "Fashion", "Makanan"}).Delete(&category.Category{}).Error
}
