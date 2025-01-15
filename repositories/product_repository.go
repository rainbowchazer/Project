package repositories

import (
	"idel/models"
	"idel/repositories/seo"

	"gorm.io/gorm"
)

// ProductRepository интерфейс для работы с продуктами
type ProductRepository interface {
	FindAllByCategoryMainIdAndItsChildren(db *gorm.DB, categoryId uint) ([]models.Product, error)
	FindAllByCategoryWhatIdAndItsChildren(db *gorm.DB, categoryId uint) ([]models.Product, error)
	FindAllByCategoryForIdAndItsChildren(db *gorm.DB, categoryId uint) ([]models.Product, error)
	FindAllByCategoryMainId(db *gorm.DB, categoryId uint) ([]models.Product, error)
	UpdateCategoryMainToNull(db *gorm.DB, categoryMainId uint) error
	UpdateCategoryWhatToNull(db *gorm.DB, categoryWhatId uint) error
	UpdateCategoryForToNull(db *gorm.DB, categoryForId uint) error
}

// productRepository реализация ProductRepository
type productRepository struct {
	seo.SeoBaseRepository[models.Product, uint]
}

// NewProductRepository конструктор для создания нового репозитория
func NewProductRepository() ProductRepository {
	return &productRepository{}
}

// FindAllByCategoryMainIdAndItsChildren находит все продукты для категории и её подкатегорий
func (r *productRepository) FindAllByCategoryMainIdAndItsChildren(db *gorm.DB, categoryId uint) ([]models.Product, error) {
	var products []models.Product
	err := db.Raw(`
		WITH RECURSIVE CategoryTree AS (
			SELECT id FROM idel_categories WHERE id = ?
			UNION ALL
			SELECT c.id FROM idel_categories c
			INNER JOIN CategoryTree ct ON ct.id = c.parent_id
		)
		SELECT p.* FROM idel_products p
		INNER JOIN CategoryTree ct ON ct.id = p.category_main_id
	`, categoryId).Scan(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// FindAllByCategoryWhatIdAndItsChildren находит все продукты для категории "what" и её подкатегорий
func (r *productRepository) FindAllByCategoryWhatIdAndItsChildren(db *gorm.DB, categoryId uint) ([]models.Product, error) {
	var products []models.Product
	err := db.Raw(`
		WITH RECURSIVE CategoryTree AS (
			SELECT id FROM idel_categories WHERE id = ?
			UNION ALL
			SELECT c.id FROM idel_categories c
			INNER JOIN CategoryTree ct ON ct.id = c.parent_id
		)
		SELECT p.* FROM idel_products p
		INNER JOIN CategoryTree ct ON ct.id = p.category_what_id
	`, categoryId).Scan(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// FindAllByCategoryForIdAndItsChildren находит все продукты для категории "for" и её подкатегорий
func (r *productRepository) FindAllByCategoryForIdAndItsChildren(db *gorm.DB, categoryId uint) ([]models.Product, error) {
	var products []models.Product
	err := db.Raw(`
		WITH RECURSIVE CategoryTree AS (
			SELECT id FROM idel_categories WHERE id = ?
			UNION ALL
			SELECT c.id FROM idel_categories c
			INNER JOIN CategoryTree ct ON ct.id = c.parent_id
		)
		SELECT p.* FROM idel_products p
		INNER JOIN CategoryTree ct ON ct.id = p.category_for_id
	`, categoryId).Scan(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// FindAllByCategoryMainId находит все продукты для основной категории
func (r *productRepository) FindAllByCategoryMainId(db *gorm.DB, categoryId uint) ([]models.Product, error) {
	var products []models.Product
	err := db.Where("category_main_id = ?", categoryId).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateCategoryMainToNull обновляет все продукты, где categoryMain равно заданному ID, и устанавливает его в NULL
func (r *productRepository) UpdateCategoryMainToNull(db *gorm.DB, categoryMainId uint) error {
	return db.Model(&models.Product{}).
		Where("category_main_id = ?", categoryMainId).
		Update("category_main_id", nil).
		Error
}

// UpdateCategoryWhatToNull обновляет все продукты, где categoryWhat равно заданному ID, и устанавливает его в NULL
func (r *productRepository) UpdateCategoryWhatToNull(db *gorm.DB, categoryWhatId uint) error {
	return db.Model(&models.Product{}).
		Where("category_what_id = ?", categoryWhatId).
		Update("category_what_id", nil).
		Error
}

// UpdateCategoryForToNull обновляет все продукты, где categoryFor равно заданному ID, и устанавливает его в NULL
func (r *productRepository) UpdateCategoryForToNull(db *gorm.DB, categoryForId uint) error {
	return db.Model(&models.Product{}).
		Where("category_for_id = ?", categoryForId).
		Update("category_for_id", nil).
		Error
}
