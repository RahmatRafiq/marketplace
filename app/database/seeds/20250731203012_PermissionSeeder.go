package seeds

import (
	"log"

	"golang_starter_kit_2025/app/features/permission"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func SeedPermissionSeeder(db *gorm.DB) error {
	log.Println("🌱 Seeding PermissionSeeder...")

	var permissions []permission.Permission
	for i := 0; i < 20; i++ {
		permissions = append(permissions, permission.Permission{
			Name:  faker.Word(),
			Group: faker.Word(),
		})
	}

	if err := db.Create(&permissions).Error; err != nil {
		return err
	}
	return nil
}
