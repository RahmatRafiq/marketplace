package seeds

import (
	"golang_starter_kit_2025/app/models"
	"log"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func SeedPermissionSeeder(db *gorm.DB) error {
	log.Println("🌱 Seeding PermissionSeeder...")

	var permissions []models.Permission
	for i := 0; i < 20; i++ {
		permissions = append(permissions, models.Permission{
			Permission: faker.Word(),
			// Group:      faker.Word(),
		})
	}

	if err := db.Create(&permissions).Error; err != nil {
		return err
	}
	return nil
}

func RollbackPermissionSeeder(db *gorm.DB) error {
	log.Println("🗑️ Rolling back PermissionSeeder…")
	return db.Unscoped().
		Where("reference LIKE ?", "PER%").
		Error
}
