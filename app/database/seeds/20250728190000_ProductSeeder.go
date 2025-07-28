package seeds

import (
	"log"
	"time"

	"golang_starter_kit_2025/app/features/product"
	"golang_starter_kit_2025/app/models"
	"math/rand"
	"strconv"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func SeedProductSeeder(db *gorm.DB) error {
	log.Println("🌱 Seeding ProductSeeder...")

	// Seeder product base
	base := models.ProductBase{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(&base).Error; err != nil {
		return err
	}

	// Seeder 20 produk varian dengan faker
	var products []product.Product
	for i := 1; i <= 20; i++ {
		name := faker.Name() + " - Varian " + string(rune('A'+i-1))
		products = append(products, product.Product{
			ProductBaseID:     base.ID,
			Name:              name,
			Slug:              faker.Username(),
			Brand:             faker.Word(),
			ShortDescription:  faker.Sentence(),
			LongDescription:   faker.Paragraph(),
			Weight:            float64(rand.Intn(400)+100) / 100.0, // 1.00 - 5.00
			Dimension1:        float64(rand.Intn(190)+10) / 100.0,  // 0.10 - 2.00
			Dimension2:        float64(rand.Intn(190)+10) / 100.0,
			Dimension3:        float64(rand.Intn(190)+10) / 100.0,
			Koli:              rand.Intn(5) + 1,
			SKU:               faker.Word() + strconv.Itoa(rand.Intn(100)),
			LowestRetailPrice: float64(rand.Intn(40001) + 10000), // 10000 - 50000
			BranchPrices:      "{}",
			Stock:             rand.Intn(100) + 1,
			Images:            []string{faker.URL(), faker.URL()},
			ReceivedAt:        time.Now(),
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		})
	}
	if err := db.Create(&products).Error; err != nil {
		return err
	}
	return nil
}

func RollbackProductSeeder(db *gorm.DB) error {
	log.Println("🗑️ Rolling back ProductSeeder…")
	return db.Unscoped().Where("name LIKE ?", "Produk Contoh - Varian %").Delete(&product.Product{}).Error
}
