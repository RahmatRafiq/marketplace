package product

import (
	"golang_starter_kit_2025/app/models"
	"golang_starter_kit_2025/facades"
	"log"

	"github.com/gin-gonic/gin"
)

type FileService struct{}

func (fs *FileService) StoreBase64File(image string, folder string, subfolder string) (*string, error) {
	return &image, nil
}

type FilterRequest struct {
	Search         *string
	OrderBy        *string
	OrderDirection *string
}

type ProductService struct {
	fileService FileService
}

func NewProductService() *ProductService {
	return &ProductService{
		fileService: FileService{},
	}
}

func (service *ProductService) GetAll(filters FilterRequest) ([]Product, error) {
	var products []Product
	query := facades.DB

	if filters.Search != nil {
		query = query.Where("name LIKE ?", "%"+*filters.Search+"%").
			Or("description LIKE ?", "%"+*filters.Search+"%").
			Or("reference LIKE ?", "%"+*filters.Search+"%")
	}

	if filters.OrderBy != nil {
		query = query.Order(*filters.OrderBy + " " + *filters.OrderDirection)
	} else {
		query = query.Order("updated_at desc")
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (service *ProductService) GetByID(id string) (Product, error) {
	var product Product
	if err := facades.DB.First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (service *ProductService) Put(ctx *gin.Context, request ProductRequest) (*Product, error) {
	var product Product

	var filenames []string
	if request.Images != nil {
		for _, image := range request.Images {
			filename, err := service.fileService.StoreBase64File(image, "images", "products")
			if err != nil {
				return nil, err
			}
			filenames = append(filenames, *filename)
			log.Println(&filename)
		}
	}

	var productBase models.ProductBase
	if request.ProductBaseID != 0 {
		facades.DB.First(&productBase, request.ProductBaseID)
	} else {
		facades.DB.Where("name = ? AND brand = ?", request.Name, request.Brand).FirstOrCreate(&productBase, models.ProductBase{})
	}

	product.ID = request.ID
	product.ProductBaseID = productBase.ID
	product.Name = request.Name
	product.Slug = request.Slug
	product.Brand = request.Brand
	product.ShortDescription = request.ShortDescription
	product.LongDescription = request.LongDescription
	product.Weight = request.Weight
	product.Dimension1 = request.Dimension1
	product.Dimension2 = request.Dimension2
	product.Dimension3 = request.Dimension3
	product.Koli = request.Koli
	product.SKU = request.SKU
	product.LowestRetailPrice = request.LowestRetailPrice
	product.BranchPrices = request.BranchPrices
	product.Stock = request.Stock
	product.Images = filenames
	product.ReceivedAt = request.ReceivedAt

	if len(request.CategoryIDs) > 0 {
		var categories []models.Category
		facades.DB.Where("id IN ?", request.CategoryIDs).Find(&categories)
		product.Categories = categories
	}
	if len(request.WarehouseIDs) > 0 {
		var warehouses []models.Warehouse
		facades.DB.Where("id IN ?", request.WarehouseIDs).Find(&warehouses)
		product.Warehouses = warehouses
	}
	if len(request.TagIDs) > 0 {
		var tags []models.Tag
		facades.DB.Where("id IN ?", request.TagIDs).Find(&tags)
		product.Tags = tags
	}
	if len(request.PromoIDs) > 0 {
		var promos []models.Promo
		facades.DB.Where("id IN ?", request.PromoIDs).Find(&promos)
		product.Promos = promos
	}
	if len(request.RelatedIDs) > 0 {
		var related []models.Product
		facades.DB.Where("id IN ?", request.RelatedIDs).Find(&related)
		product.Related = related
	}

	if count := facades.DB.Model(&Product{}).Where("id = ?", request.ID).Find(&map[string]interface{}{}).RowsAffected; count == 0 {
		if err := facades.DB.Create(&product).Error; err != nil {
			return &product, err
		}
	} else {
		if err := facades.DB.Model(&Product{}).Where("id = ?", request.ID).Updates(&product).Error; err != nil {
			return &product, err
		}
		if err := facades.DB.First(&product, request.ID).Error; err != nil {
			return &product, err
		}
	}

	return &product, nil
}

func (service *ProductService) Delete(id string) error {
	result := facades.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}
