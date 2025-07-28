package permission

import (
	"golang_starter_kit_2025/facades"
)

type PermissionService struct{}

func NewPermissionService() *PermissionService {
	return &PermissionService{}
}

func (*PermissionService) GetAll() ([]Permission, error) {
	var permissions []Permission
	if err := facades.DB.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (*PermissionService) Put(updatedPermission Permission) (Permission, error) {
	var permission Permission

	if count := facades.DB.Model(&Permission{}).Where("id = ?", updatedPermission.ID).Find(&map[string]interface{}{}).RowsAffected; count == 0 {
		if err := facades.DB.Create(&updatedPermission).Error; err != nil {
			return permission, err
		}
	} else {
		if err := facades.DB.Where("id = ?", updatedPermission.ID).Updates(&updatedPermission).Error; err != nil {
			return permission, err
		}

		if err := facades.DB.First(&permission, updatedPermission.ID).Error; err != nil {
			return permission, err
		}
	}

	return permission, nil
}

func (*PermissionService) Delete(id string) error {
	var permission Permission
	if err := facades.DB.First(&permission, id).Error; err != nil {
		return err
	}
	return facades.DB.Delete(&permission).Error
}
