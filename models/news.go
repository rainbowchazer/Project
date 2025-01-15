package models

import (
	"idel/models/seo"
	"time"
)

type News struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Seo       seo.Seo   `gorm:"embedded" json:"seo"`
	Title     string    `gorm:"not null" json:"title"`
	Date      time.Time `gorm:"not null" json:"date"`
	Text      string    `gorm:"not null" json:"text"`
	ImagePath *string   `gorm:"column:image_path" json:"image_path,omitempty"`
}

func (News) TableName() string {
	return "idel_news"
}

func NewNews(title string, text string, image_path *string) *News {
	return &News{
		Title:     title,
		Text:      text,
		ImagePath: image_path,
		Date:      time.Now(),
	}
}

func (n *News) GetSeoIdSource() string {
	return n.Title
}

func (n *News) GetId() any {
	return n.ID
}

func (n *News) GetSeo() seo.Seo {
	return n.Seo
}

func (n *News) SetSeo(s seo.Seo) {
	n.Seo = s
}
