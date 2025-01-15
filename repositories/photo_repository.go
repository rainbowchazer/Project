package repositories

import (
	"idel/models"
	"idel/repositories/seo"
)

// NewsRepository - интерфейс для работы с новостями
type PhotoRepository interface {
	seo.SeoBaseRepository[models.Photo, uint] // Встраиваем SeoBaseRepository
}

// newsRepository - структура, реализующая интерфейс NewsRepository
type photoRepository struct {
	seo.SeoBaseRepository[models.Photo, uint]
}

// NewNewsRepository - конструктор для newsRepository
func NewPhotoRepository() PhotoRepository {
	return &photoRepository{
		SeoBaseRepository: seo.NewSeoBaseRepository[models.Photo, uint](),
	}
}
