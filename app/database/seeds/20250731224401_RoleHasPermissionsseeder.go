package seeds

import (
	"fmt"
	"golang_starter_kit_2025/app/models"
	"log"

	"gorm.io/gorm"
)

func SeedRoleHasPermissionsSeeder(db *gorm.DB) error {
	log.Println("🌱 Seeding RoleHasPermissionsSeeder...")

	var roles []models.Role
	var permissions []models.Permission

	if err := db.Find(&roles).Error; err != nil {
		return err
	}
	if len(roles) == 0 {
		return fmt.Errorf("Seeder gagal: data role kosong. Jalankan seeder Role terlebih dahulu.")
	}
	if err := db.Find(&permissions).Error; err != nil {
		return err
	}
	if len(permissions) == 0 {
		return fmt.Errorf("Seeder gagal: data permission kosong. Jalankan seeder Permission terlebih dahulu.")
	}

	var roleHasPermissions []models.RoleHasPermissions
	for _, role := range roles {
		for _, permission := range permissions {
			roleHasPermissions = append(roleHasPermissions, models.RoleHasPermissions{
				RoleID:       role.ID,
				PermissionID: permission.ID,
			})
		}
	}

	if err := db.Create(&roleHasPermissions).Error; err != nil {
		return err
	}
	return nil
}

func RollbackRoleHasPermissionsSeeder(db *gorm.DB) error {
	log.Println("🗑️ Rolling back RoleHasPermissionsSeeder…")
	return db.Unscoped().Delete(&models.RoleHasPermissions{}).Error
}
