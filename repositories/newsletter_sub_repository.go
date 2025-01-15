package repositories

import (
	"idel/models"

	"gorm.io/gorm"
)

// NewsletterSubscriberRepository - интерфейс для работы с подписчиками на рассылку
type NewsletterSubscriberRepository interface {
	FindAllEmails(db *gorm.DB) ([]string, error)
	FindByEmail(db *gorm.DB, email string) (*models.NewsletterSubscriber, error)
	Create(db *gorm.DB, subscriber *models.NewsletterSubscriber) error
}

// newsletterSubscriberRepository - структура, реализующая NewsletterSubscriberRepository
type newsletterSubscriberRepository struct{}

// NewNewsletterSubscriberRepository - конструктор для newsletterSubscriberRepository
func NewNewsletterSubscriberRepository() NewsletterSubscriberRepository {
	return &newsletterSubscriberRepository{}
}

// FindAllEmails - получение всех email адресов подписчиков
func (r *newsletterSubscriberRepository) FindAllEmails(db *gorm.DB) ([]string, error) {
	var emails []string
	err := db.Model(&models.NewsletterSubscriber{}).
		Select("email").
		Find(&emails).Error
	return emails, err
}

// FindByEmail - поиск подписчика по email
func (r *newsletterSubscriberRepository) FindByEmail(db *gorm.DB, email string) (*models.NewsletterSubscriber, error) {
	var subscriber models.NewsletterSubscriber
	err := db.Where("email = ?", email).First(&subscriber).Error
	if err != nil {
		return nil, err
	}
	return &subscriber, nil
}

// Create - создание нового подписчика
func (r *newsletterSubscriberRepository) Create(db *gorm.DB, subscriber *models.NewsletterSubscriber) error {
	return db.Create(subscriber).Error
}
