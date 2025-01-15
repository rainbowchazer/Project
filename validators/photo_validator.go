package validators

import (
	"fmt"
	"idel/dtos"
	"idel/models"
	"idel/repositories"

	"gorm.io/gorm"
)

type PhotoValidator struct {
	repo repositories.PhotoRepository
	db   *gorm.DB
}

func NewPhotoValidator(repo repositories.PhotoRepository, db *gorm.DB) *PhotoValidator {
	return &PhotoValidator{
		repo: repo,
		db:   db,
	}
}

func (p *PhotoValidator) ThrowIfNotExists(id uint) (*models.Photo, error) {
	var photo models.Photo
	err := p.db.First(&photo, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("photo not found with id:%d", id)
		}
		return nil, fmt.Errorf("database error while checking existence of photo id:%d", id)
	}
	return &photo, nil
}

func (p *PhotoValidator) ThrowIfExists(id uint) error {
	var count int64
	err := p.db.Model(&models.Photo{}).Where("id == ?", id).Count(&count).Error
	if err != nil {
		return fmt.Errorf("database error while checking existence of photo id:%d", id)
	}
	if count > 0 {
		return fmt.Errorf("photo already exists with id: %d", id)
	}
	return nil
}

func (p *PhotoValidator) ValidateImage(photoDto dtos.PhotoDto) error {
	if (photoDto.Image == nil) || (photoDto.Image.Size == 0) && len(photoDto.ImagePath) == 0 {
		return fmt.Errorf("photo dto must contain either an image file or an image path")
	}
	return nil
}

func (p *PhotoValidator) ValidateToUpdate(photoDto dtos.PhotoDto) (*models.Photo, error) {
	photo, err := p.ThrowIfNotExists(uint(photoDto.ID))
	if err != nil {
		return nil, err
	}

	if err := p.ValidateImage(photoDto); err != nil {
		return nil, err
	}

	return photo, nil
}
