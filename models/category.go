package models

import (
	"idel/models/seo"
)

// Category представляет категорию с SEO и иерархической структурой.
type Category struct {
	ID           uint    `gorm:"primaryKey;autoIncrement" json:"id"`            // Первичный ключ
	Seo          seo.Seo `gorm:"embedded" json:"seo"`                           // Встроенная структура для SEO
	Name         string  `gorm:"not null" json:"name"`                          // Название категории
	Description  *string `json:"description,omitempty"`                         // Описание (может быть NULL)
	ShowInSlider bool    `gorm:"column:show_in_slider" json:"show_in_slider"`   // Отображение в слайдере
	ImagePath    *string `gorm:"column:image_path" json:"image_path,omitempty"` // Путь к изображению (может быть NULL)

	// Связь "многие к одному" с указанием внешнего ключа
	ParentID *uint     `gorm:"column:parent_id" json:"parent_id,omitempty"` // ID родителя
	Parent   *Category `gorm:"foreignKey:ParentID" json:"-"`                // Родительская категория

	// Связь "один ко многим"
	Children []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"` // Дочерние категории

	// Виртуальное поле, не сохраняется в базу данных
	VirtualParentID *uint `gorm:"-" json:"virtual_parent_id,omitempty"`
}

func (Category) TableName() string {
	return "idel_categories"
}

// Реализация метода для получения ParentID
func (c *Category) GetParentID() *uint {
	return c.ParentID
}

// Реализация интерфейса seo.ISeoEntity

// GetSeoIdSource возвращает имя категории как источник для SEO ID
func (c *Category) GetSeoIdSource() string {
	return c.Name
}

// GetId возвращает ID категории
func (c *Category) GetId() any {
	return c.ID
}

// GetSeo возвращает структуру Seo
func (c *Category) GetSeo() seo.Seo {
	return c.Seo
}

// SetSeo устанавливает структуру Seo
func (c *Category) SetSeo(s seo.Seo) {
	c.Seo = s
}
