package seeds

import (
	"golang_starter_kit_2025/app/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func SeedProductSeeder(db *gorm.DB) error {
	log.Println("🌱 Seeding ProductSeeder...")

	// Contoh seeder product base
	base := models.ProductBase{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(&base).Error; err != nil {
		return err
	}

	// Contoh seeder produk varian
	products := []models.Product{
		{
			ProductBaseID:     base.ID,
			Name:              "Produk Contoh - Varian A",
			Slug:              "produk-contoh-varian-a",
			Brand:             "BrandA",
			ShortDescription:  "Deskripsi pendek varian A",
			LongDescription:   "Deskripsi panjang varian A",
			Weight:            1.2,
			Dimension1:        1.0,
			Dimension2:        0.5,
			Dimension3:        0.3,
			Koli:              1,
			SKU:               "SKU-A",
			LowestRetailPrice: 10000,
			BranchPrices:      "{}",
			Stock:             10,
			Images:            []string{"imageA1.jpg", "imageA2.jpg"},
			ReceivedAt:        time.Now(),
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		{
			ProductBaseID:     base.ID,
			Name:              "Produk Contoh - Varian B",
			Slug:              "produk-contoh-varian-b",
			Brand:             "BrandA",
			ShortDescription:  "Deskripsi pendek varian B",
			LongDescription:   "Deskripsi panjang varian B",
			Weight:            1.3,
			Dimension1:        1.1,
			Dimension2:        0.6,
			Dimension3:        0.4,
			Koli:              1,
			SKU:               "SKU-B",
			LowestRetailPrice: 12000,
			BranchPrices:      "{}",
			Stock:             8,
			Images:            []string{"imageB1.jpg"},
			ReceivedAt:        time.Now(),
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
	}
	if err := db.Create(&products).Error; err != nil {
		return err
	}
	return nil
}

func RollbackProductSeeder(db *gorm.DB) error {
	log.Println("🗑️ Rolling back ProductSeeder…")
	return db.Unscoped().Where("name LIKE ?", "Produk Contoh - Varian %").Delete(&models.Product{}).Error
}
