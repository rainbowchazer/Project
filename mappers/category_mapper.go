package mappers

import (
	"idel/dtos"
	"idel/models"
)

type CategoryMapper struct{}

func (m *CategoryMapper) MapToEntity(source *dtos.CategoryDto) *models.Category {
	if source == nil {
		return nil
	}

	return &models.Category{
		ID:           source.ID,
		Seo:          source.Seo,
		ParentID:     &source.ID,
		ShowInSlider: source.ShowInSlider,
		Name:         source.Name,
		Description:  &source.Description,
		ImagePath:    &source.ImagePath,
	}
}

func (m *CategoryMapper) MapToDto(source *models.Category) *dtos.CategoryDto {
	if source == nil {
		return nil
	}

	dto := &dtos.CategoryDto{
		ID:           source.ID,
		Seo:          source.Seo,
		Name:         source.Name,
		ShowInSlider: source.ShowInSlider,
	}

	if source.Description != nil {
		dto.Description = *source.Description
	}

	if source.ImagePath != nil {
		dto.ImagePath = *source.ImagePath
	}

	if source.ParentID != nil {
		dto.ParentID = source.ParentID
	}

	return dto
}
