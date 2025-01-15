package seo

type SeoMetaTag struct {
	Name    string `gorm:"column:name"`
	Content string `gorm:"column:content"`
}
