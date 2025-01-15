package seo

import (
	"strings"

	"gorm.io/gorm"
)

// SeoBaseRepository - интерфейс для базовых операций с Seo
type SeoBaseRepository[T any, ID comparable] interface {
	FindBySeoId(db *gorm.DB, seoId string) (*T, error)
	CountBySeoIdIgnoreCase(db *gorm.DB, seoId string) (int64, error)
	CountBySeoIdIgnoreCaseAndIdNot(db *gorm.DB, seoId string, excludeId ID) (int64, error)
	IsExistBySeoId(db *gorm.DB, seoId string, excludeId *ID) (bool, error)
}

// seoBaseRepository - структура, реализующая интерфейс SeoBaseRepository
type seoBaseRepository[T any, ID comparable] struct{}

// NewSeoBaseRepository - конструктор для seoBaseRepository
func NewSeoBaseRepository[T any, ID comparable]() SeoBaseRepository[T, ID] {
	return &seoBaseRepository[T, ID]{}
}

// FindBySeoId - поиск записи по SeoId
func (r *seoBaseRepository[T, ID]) FindBySeoId(db *gorm.DB, seoId string) (*T, error) {
	var entity T
	err := db.Where("LOWER(seo_id) = ?", strings.ToLower(seoId)).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// CountBySeoIdIgnoreCase - подсчёт записей с SeoId (регистронезависимо)
func (r *seoBaseRepository[T, ID]) CountBySeoIdIgnoreCase(db *gorm.DB, seoId string) (int64, error) {
	var count int64
	err := db.Model(new(T)).Where("LOWER(seo_id) = ?", strings.ToLower(seoId)).Count(&count).Error
	return count, err
}

// CountBySeoIdIgnoreCaseAndIdNot - подсчёт записей с SeoId, исключая запись с определённым ID
func (r *seoBaseRepository[T, ID]) CountBySeoIdIgnoreCaseAndIdNot(db *gorm.DB, seoId string, excludeId ID) (int64, error) {
	var count int64
	err := db.Model(new(T)).
		Where("LOWER(seo_id) = ? AND id != ?", strings.ToLower(seoId), excludeId).
		Count(&count).Error
	return count, err
}

// IsExistBySeoId - проверка существования записи с SeoId
func (r *seoBaseRepository[T, ID]) IsExistBySeoId(db *gorm.DB, seoId string, excludeId *ID) (bool, error) {
	if excludeId != nil {
		count, err := r.CountBySeoIdIgnoreCaseAndIdNot(db, seoId, *excludeId)
		return count > 0, err
	}
	count, err := r.CountBySeoIdIgnoreCase(db, seoId)
	return count > 0, err
}
