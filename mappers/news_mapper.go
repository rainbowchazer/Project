package mappers

import (
	"idel/dtos"
	"idel/models"
)

type NewsMapper struct{}

func (m *NewsMapper) MapToEntity(source *dtos.NewsDto) *models.News {
	if source == nil {
		return nil
	}

	return &models.News{
		ID:        source.Id,
		Seo:       source.Seo,
		Title:     source.Title,
		Date:      source.Date,
		Text:      source.Text,
		ImagePath: &source.ImagePath,
	}
}

func (m *NewsMapper) MapToDto(source *models.News) *dtos.NewsDto {
	if source == nil {
		return nil
	}

	imagePath := ""
	if source.ImagePath != nil {
		imagePath = *source.ImagePath
	}

	return &dtos.NewsDto{
		Id:        source.ID,
		Seo:       source.Seo,
		Title:     source.Title,
		Date:      source.Date,
		Text:      source.Text,
		ImagePath: imagePath,
	}
}
