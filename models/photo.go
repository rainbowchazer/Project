package models

type Photo struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Description *string `json:"description"`
	ImagePath   string  `gorm:"column:image_path;not null" json:"image_path,omitempty"`
}

func (Photo) TableName() string {
	return "idel_photos"
}
