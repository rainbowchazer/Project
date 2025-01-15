package validators

import (
	"fmt"
	"idel/models"
	"idel/repositories"

	"gorm.io/gorm"
)

type CategoryValidator struct {
	repo repositories.CategoryRepository
	db   *gorm.DB
}

func NewCategoryValidator(repo repositories.CategoryRepository, db *gorm.DB) *CategoryValidator {
	return &CategoryValidator{
		repo: repo,
		db:   db,
	}
}

func (v *CategoryValidator) ThrowIfNotExists(id uint) (*models.Category, error) {
	var category models.Category
	err := v.db.First(&category, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("category not found with ID: %d", id)
		}
		return nil, fmt.Errorf("database error while checking existence of category with ID: %d, error: %v", id, err)
	}
	return &category, nil
}

func (v *CategoryValidator) ThrowIfExists(id uint) error {
	if id == 0 {
		return nil
	}
	var count int64
	err := v.db.Model(&models.Category{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return fmt.Errorf("database error while checking existence of category with ID: %d, error: %v", id, err)
	}
	if count > 0 {
		return fmt.Errorf("category already exists with ID:%d", id)
	}
	return nil
}

func (v *CategoryValidator) ThrowIfNotUnique(categoryName string) error {
	count, err := v.repo.CountByNameIgnoreCase(v.db, categoryName)
	if err != nil {
		return fmt.Errorf("database error while checking uniqueness of category name: %s", categoryName)
	}
	if count > 0 {
		return fmt.Errorf("category already exists with name: %s", categoryName)
	}
	return nil
}

func (v *CategoryValidator) ThrowIfNotUniqueExceptId(id uint, categoryName string) error {
	count, err := v.repo.CountByNameIgnoreCaseAndNotId(v.db, categoryName, id)
	if err != nil {
		return fmt.Errorf("database error while checking uniqueness of category name: %s", categoryName)
	}
	if count > 0 {
		return fmt.Errorf("category already exists with name: %s excluding ID: %d", categoryName, id)
	}
	return nil
}
