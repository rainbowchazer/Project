package repositories

import (
	"idel/models"
	"idel/repositories/seo"
)

// NewsRepository - интерфейс для работы с новостями
type NewsRepository interface {
	seo.SeoBaseRepository[models.News, uint] // Встраиваем SeoBaseRepository
}

// newsRepository - структура, реализующая интерфейс NewsRepository
type newsRepository struct {
	seo.SeoBaseRepository[models.News, uint]
}

// NewNewsRepository - конструктор для newsRepository
func NewNewsRepository() NewsRepository {
	return &newsRepository{
		SeoBaseRepository: seo.NewSeoBaseRepository[models.News, uint](),
	}
}
