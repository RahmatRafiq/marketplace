package seeds

import (
	"golang_starter_kit_2025/app/models"
	"log"

	"gorm.io/gorm"
)

func SeedPermissionSeeder(db *gorm.DB) error {
	log.Println("🌱 Seeding PermissionSeeder...")

	permissionNames := []string{
		"lihat_produk", "tambah_produk", "edit_produk", "hapus_produk", "lihat_kategori", "tambah_kategori", "edit_kategori", "hapus_kategori", "lihat_user", "tambah_user", "edit_user", "hapus_user", "akses_laporan", "export_data", "import_data", "verifikasi_transaksi", "ubah_status", "akses_diskon", "akses_promo", "akses_stok",
	}
	var permissions []models.Permission
	for _, name := range permissionNames {
		permissions = append(permissions, models.Permission{
			Permission: name,
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
		Where("permission LIKE ?", "%").
		Delete(&models.Permission{}).
		Error
}
