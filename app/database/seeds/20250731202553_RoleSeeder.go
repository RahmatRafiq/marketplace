package seeds

import (
	"golang_starter_kit_2025/app/models"
	"log"

	"gorm.io/gorm"
)

func SeedRoleSeeder(db *gorm.DB) error {
	log.Println("🌱 Seeding RoleSeeder...")

	data := models.Role{
		Role: "User",
	}
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func RollbackRoleSeeder(db *gorm.DB) error {
	log.Println("🗑️ Rolling back RoleSeeder…")
	return db.Unscoped().
		Where("reference LIKE ?", "USR%").
		Error
}
