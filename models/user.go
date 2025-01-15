package models

import "idel/enums"

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`

	Role                 enums.UserRole   `json:"role"`
	Status               enums.UserStatus `json:"status"`
	SubscribedOnFeedback bool             `gorm:"column:subscribed_on_feedback;not null" json:"subscribed_on_feedback"`
}

func (User) TableName() string {
	return "idel_users"
}

func NewUser(email string, password string, role enums.UserRole, status enums.UserStatus) *User {
	return &User{
		Email:    email,
		Password: password,
		Role:     role,
		Status:   status,
	}
}
