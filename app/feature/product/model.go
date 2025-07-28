package product

import (
	"golang_starter_kit_2025/app/helpers"
	"golang_starter_kit_2025/app/models"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                uint                `gorm:"primaryKey" json:"id"`
	Reference         string              `json:"reference"`
	ProductBaseID     uint                `json:"product_base_id"`
	ProductBase       *models.ProductBase `json:"product_base"`
	Name              string              `json:"name"`
	Slug              string              `json:"slug"`
	Brand             string              `json:"brand"`
	ShortDescription  string              `json:"short_description"`
	LongDescription   string              `json:"long_description"`
	Weight            float64             `json:"weight"`
	Dimension1        float64             `json:"dimension_1" gorm:"column:dimension_1"`
	Dimension2        float64             `json:"dimension_2" gorm:"column:dimension_2"`
	Dimension3        float64             `json:"dimension_3" gorm:"column:dimension_3"`
	Koli              int                 `json:"koli"`
	SKU               string              `json:"sku"`
	LowestRetailPrice float64             `json:"lowest_retail_price"`
	BranchPrices      string              `json:"branch_prices" gorm:"type:json"`
	Stock             int                 `json:"stock"`
	Images            []string            `json:"images" gorm:"serializer:json"`
	ReceivedAt        time.Time           `json:"received_at"`

	Categories []models.Category  `gorm:"many2many:product_categories;" json:"categories"`
	Warehouses []models.Warehouse `gorm:"many2many:product_warehouses;" json:"warehouses"`
	Related    []models.Product   `gorm:"many2many:product_related;association_jointable_foreignkey:related_product_id" json:"related_products"`
	Tags       []models.Tag       `gorm:"many2many:product_tags;" json:"tags"`
	Promos     []models.Promo     `gorm:"many2many:product_promos;" json:"promos"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}

// BeforeCreate hook
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.Reference = helpers.GenerateReference("PRD")
	return
}

// AfterFind hook
func (m *Product) AfterFind(tx *gorm.DB) (err error) {
	if m.Images != nil && len(m.Images) > 0 {
		for i, Image := range m.Images {
			m.Images[i] = helpers.GetFileURL(Image, "member_lands")
		}
	}
	return
}

// AfterCreate hook
func (m *Product) AfterCreate(tx *gorm.DB) (err error) {
	if m.Images != nil && len(m.Images) > 0 {
		for i, image := range m.Images {
			m.Images[i] = helpers.GetFileURL(image, "member_lands")
		}
	}
	return
}

// AfterUpdate hook
func (m *Product) AfterUpdate(tx *gorm.DB) (err error) {
	if m.Images != nil && len(m.Images) > 0 {
		for i, image := range m.Images {
			m.Images[i] = helpers.GetFileURL(image, "member_lands")
		}
	}
	return
}
