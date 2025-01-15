package repositories

import (
	"idel/models"

	"gorm.io/gorm"
)

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	FindByEmailWithStatusActive(email string) (*models.User, error)
	FindAllAdmins() ([]models.User, error)
	FindAllFeedbackSubscribers() ([]string, error)
}

// userRepository реализация UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository конструктор для создания нового репозитория
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// FindByEmail находит пользователя по email
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmailWithStatusActive находит пользователя по email и статусу "ACTIVE"
func (r *userRepository) FindByEmailWithStatusActive(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? AND status = ?", email, "ACTIVE").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAllAdmins находит всех пользователей с ролью "ADMIN"
func (r *userRepository) FindAllAdmins() ([]models.User, error) {
	var users []models.User
	err := r.db.Where("role = ?", "ADMIN").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// FindAllFeedbackSubscribers находит все email-адреса пользователей, которые подписаны на обратную связь
func (r *userRepository) FindAllFeedbackSubscribers() ([]string, error) {
	var emails []string
	err := r.db.Model(&models.User{}).Where("subscribed_on_feedback = ?", true).Pluck("email", &emails).Error
	if err != nil {
		return nil, err
	}
	return emails, nil
}
