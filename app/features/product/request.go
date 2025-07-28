package product

import "time"

type ProductRequest struct {
	ID                uint      `json:"id,omitempty" form:"id,omitempty"`
	ProductBaseID     uint      `json:"product_base_id" form:"product_base_id"`
	Name              string    `json:"name" form:"name" binding:"required"`
	Slug              string    `json:"slug" form:"slug"`
	Brand             string    `json:"brand" form:"brand"`
	ShortDescription  string    `json:"short_description" form:"short_description"`
	LongDescription   string    `json:"long_description" form:"long_description"`
	Weight            float64   `json:"weight" form:"weight"`
	Dimension1        float64   `json:"dimension_1" form:"dimension_1"`
	Dimension2        float64   `json:"dimension_2" form:"dimension_2"`
	Dimension3        float64   `json:"dimension_3" form:"dimension_3"`
	Koli              int       `json:"koli" form:"koli"`
	SKU               string    `json:"sku" form:"sku"`
	LowestRetailPrice float64   `json:"lowest_retail_price" form:"lowest_retail_price"`
	BranchPrices      string    `json:"branch_prices" form:"branch_prices"`
	Stock             int       `json:"stock" form:"stock"`
	Images            []string  `json:"images" form:"images" type:"array:string"`
	ReceivedAt        time.Time `json:"received_at" form:"received_at"`
	CategoryIDs       []uint    `json:"category_ids" form:"category_ids"`
	WarehouseIDs      []uint    `json:"warehouse_ids" form:"warehouse_ids"`
	TagIDs            []uint    `json:"tag_ids" form:"tag_ids"`
	PromoIDs          []uint    `json:"promo_ids" form:"promo_ids"`
	RelatedIDs        []uint    `json:"related_ids" form:"related_ids"`
}
