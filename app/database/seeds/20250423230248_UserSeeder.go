package seeds

import (
	"golang_starter_kit_2025/app/features/user"
	"golang_starter_kit_2025/app/helpers"
	"log"
	"time"

	"gorm.io/gorm"
)

func SeedUserSeeder(db *gorm.DB) error {
	log.Println("🌱 Seeding UserSeeder...")

	data := user.User{
		Reference: helpers.GenerateReference("USR"),
		Username:  "admin",
		Email:     "admin@example.com",
		Password:  "admin@example.com",
		Pin:       "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
func RollbackUserSeeder(db *gorm.DB) error {
	log.Println("🗑️ Rolling back UserSeeder…")
	return db.Unscoped().
		Where("username = ?", "admin").
		Delete(&user.User{}).
		Error
}
