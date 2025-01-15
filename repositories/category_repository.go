package repositories

import (
	"idel/models"
	"idel/repositories/seo"
	"strings"

	"gorm.io/gorm"
)

// CategoryRepository - интерфейс для работы с категориями
type CategoryRepository interface {
	seo.SeoBaseRepository[models.Category, uint]
	CountByNameIgnoreCase(db *gorm.DB, categoryName string) (int64, error)
	CountByNameIgnoreCaseAndNotId(db *gorm.DB, categoryName string, excludeId uint) (int64, error)
	FindAllParentCategories(db *gorm.DB) ([]models.Category, error)
}

// categoryRepository - структура, реализующая CategoryRepository
type categoryRepository struct {
	seo.SeoBaseRepository[models.Category, uint]
}

// NewCategoryRepository - конструктор для categoryRepository
func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		SeoBaseRepository: seo.NewSeoBaseRepository[models.Category, uint](),
	}
}

// CountByNameIgnoreCase - подсчёт категорий с именем без учёта регистра
func (r *categoryRepository) CountByNameIgnoreCase(db *gorm.DB, categoryName string) (int64, error) {
	var count int64
	err := db.Model(&models.Category{}).
		Where("LOWER(name) = ?", strings.ToLower(categoryName)).
		Count(&count).Error
	return count, err
}

// CountByNameIgnoreCaseAndNotId - подсчёт категорий с именем без учёта регистра, исключая запись с определённым ID
func (r *categoryRepository) CountByNameIgnoreCaseAndNotId(db *gorm.DB, categoryName string, excludeId uint) (int64, error) {
	var count int64
	err := db.Model(&models.Category{}).
		Where("LOWER(name) = ? AND id != ?", strings.ToLower(categoryName), excludeId).
		Count(&count).Error
	return count, err
}

// FindAllParentCategories - получение всех родительских категорий
func (r *categoryRepository) FindAllParentCategories(db *gorm.DB) ([]models.Category, error) {
	var categories []models.Category
	err := db.Where("parent_id IS NULL OR parent_id = 0").Find(&categories).Error
	return categories, err
}
