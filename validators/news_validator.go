package validators

import (
	"fmt"
	"idel/models"
	"idel/repositories"

	"gorm.io/gorm"
)

type NewsValidator struct {
	repo repositories.NewsRepository
	db   *gorm.DB
}

func NewNewsValidator(repo repositories.NewsRepository, db *gorm.DB) *NewsValidator {
	return &NewsValidator{
		repo: repo,
		db:   db,
	}
}

func (n *NewsValidator) ThrowIfNotExists(id uint) (*models.News, error) {
	var news models.News
	err := n.db.First(&news, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("news is not found")
		}
		return nil, fmt.Errorf("database error while checking existence of news id: %d", id)
	}
	return &news, nil
}

func (n *NewsValidator) ThrowIfExists(id uint) error {
	var count int64
	err := n.db.Model(&models.News{}).Where("id == ?", id).Count(&count).Error
	if err != nil {
		return fmt.Errorf("database error while checking existence of news id: %d", id)
	}
	if count > 0 {
		return fmt.Errorf("news already exists with id: %d", id)
	}
	return nil
}
