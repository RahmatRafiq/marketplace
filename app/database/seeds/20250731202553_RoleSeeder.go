package seeds

import (
	"golang_starter_kit_2025/app/models"
	"log"

	"gorm.io/gorm"
)

func SeedRoleSeeder(db *gorm.DB) error {
	log.Println("🌱 Seeding RoleSeeder...")

	roleNames := []string{"Admin", "User", "Manajer", "Kasir", "Supervisor", "Staf", "Pelanggan", "Teknisi", "Marketing", "Auditor"}
	var roles []models.Role
	for _, name := range roleNames {
		roles = append(roles, models.Role{
			Role: name,
		})
	}
	if err := db.Create(&roles).Error; err != nil {
		return err
	}
	return nil
}

func RollbackRoleSeeder(db *gorm.DB) error {
	log.Println("🗑️ Rolling back RoleSeeder…")
	return db.Unscoped().
		Where("role LIKE ?", "%").
		Delete(&models.Role{}).
		Error
}
