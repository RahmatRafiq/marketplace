package category

import (
	"golang_starter_kit_2025/facades"

	"gorm.io/gorm/clause"
)

type CategoryService struct{}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) GetAll() ([]Category, error) {
	var categories []Category
	if err := facades.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CategoryService) GetByID(id string) (Category, error) {
	var category Category
	if err := facades.DB.First(&category, id).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (s *CategoryService) Put(category Category) (Category, error) {
	if err := facades.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"category", "updated_at"}),
	}).Create(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (s *CategoryService) Delete(id string) error {
	var category Category
	if err := facades.DB.First(&category, id).Error; err != nil {
		return err
	}
	return facades.DB.Delete(&category).Error
}
